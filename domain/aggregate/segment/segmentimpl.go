package segment

import (
	"context"
	"github.com/JerryZhou343/leaf-go/domain/aggregate/segment/entity"
	"github.com/JerryZhou343/leaf-go/domain/util"
	leaf_go "github.com/JerryZhou343/leaf-go/genproto"
	"github.com/bilibili/kratos/pkg/log"
	"sync"
	"time"
)

const (
	MAX_STEP         = 1000000
	SEGMENT_DURATION = 15 * time.Minute
)
//[string]*entity.SegmentBuffer
type SegmentImpl struct {
	cache  sync.Map
	repo   Repo
	initOK bool
}

func NewSegmentImpl(repo Repo) *SegmentImpl {
	ret := &SegmentImpl{
		repo: repo,
		cache: sync.Map{},
	}
	ret.initOK = false
	return ret
}

func (s *SegmentImpl) Init() (err error) {
	err = s.updateCacheFromDB()
	if err != nil {
		return
	}
	s.updateCacheFromDBAtEveryMinute()
	s.initOK = true
	return nil
}

func (s *SegmentImpl) updateCacheFromDB() (err error) {
	var (
		dbTags []string
	)
	dbTags, err = s.repo.GetAllTags(context.Background())
	if err != nil {
		return
	}

	if dbTags == nil && len(dbTags) == 0 {
		return
	}

	cacheTagsSet := map[string]struct{}{}
	insertTagsSet := map[string]struct{}{}
	removeTagsSet := map[string]struct{}{}
	//待插入tag
	for _, tag := range dbTags {
		insertTagsSet[tag] = struct{}{}
	}

	//当前已经缓存tag
	s.cache.Range(func(key, value interface{}) bool {
		cacheTagsSet[key.(string)] = struct{}{}
		removeTagsSet[key.(string)] = struct{}{}
		return true
	})


	//插入新tag
	//1) 移除已经包含的，留下待添加到cache的
	for k, _ := range cacheTagsSet {
		if _, ok := insertTagsSet[k]; ok {
			delete(insertTagsSet, k)
		}
	}
	//2) 将新增的插入到cache中
	for k, _ := range insertTagsSet {
		buffer := entity.NewSegmentBuffer(k)
		segment := buffer.GetCurrent()
		segment.Value.Store(0)
		segment.Max.Store(0)
		segment.Step.Store(0)
		s.cache.Store(k,buffer)
	}

	//cache中已经失效的tags从cache中删除
	//1)过滤出需要删除的tag
	for _, tag := range dbTags {
		if _, ok := removeTagsSet[tag]; ok {
			delete(removeTagsSet, tag)
		}
	}
	//2)remove
	for tag, _ := range removeTagsSet {
		s.cache.Delete(tag)
	}
	return nil
}

func (s *SegmentImpl) updateCacheFromDBAtEveryMinute() {
	go func() {
		for {
			select {
			case <-time.After(1 * time.Minute):
				s.updateCacheFromDB()
			}
		}
	}()
}

func (s *SegmentImpl) Get(ctx context.Context, key string) (int64, error) {
	if !s.initOK {
		return 0, leaf_go.ID_ID_CACHE_INIT_FALSE
	}
	if buffer, ok := s.cache.Load(key); ok {
		if !buffer.(*entity.SegmentBuffer).IsInitOk() {
			func(buffer *entity.SegmentBuffer) {
				buffer.Lock.Lock()
				defer buffer.Lock.Unlock()
				if !buffer.IsInitOk() {
					err := s.updateSegmentFromDB(ctx, key, buffer.GetCurrent())
					if err == nil {
						buffer.SetInitOK(true)
					}
				}
			}(buffer.(*entity.SegmentBuffer))

		}
		buffer,_ := s.cache.Load(key)
		return s.getIdFromSegmentBuffer(ctx, buffer.(*entity.SegmentBuffer))
	}
	return 0, leaf_go.ID_KEY_NOT_EXISTS
}

func (s *SegmentImpl) getIdFromSegmentBuffer(ctx context.Context, buffer *entity.SegmentBuffer) (id int64, err error) {
	for {
		buffer.Lock.RLock()
		//获得当前的segment
		segment := buffer.GetCurrent()
		//检查备用segment状况
		if !buffer.NextReady && //没有就绪
			segment.GetIdle() < int64(0.9*float32(segment.Step.Load())) && //当前空闲不够
			buffer.ThreadRunning.CAS(false, true) { //切换线程运行状态
			go func() {
				next := buffer.Segments[buffer.NextPos()]
				updateErr := s.updateSegmentFromDB(ctx, buffer.Key, next)
				if updateErr == nil {
					buffer.Lock.Lock()
					buffer.NextReady = true
					buffer.Lock.Unlock()
				}
				buffer.ThreadRunning.Store(false)
			}()
		}

		//尝试第一次拿id
		id = segment.Incr()
		if id < segment.GetMax() {
			buffer.Lock.RUnlock()
			return
		}
		buffer.Lock.RUnlock()
		//sleep
		s.waitAndSleep(buffer)

		//尝试第二次拿id
		buffer.Lock.Lock()
		segment = buffer.GetCurrent()
		id = segment.Incr()
		if id < segment.GetMax() {
			buffer.Lock.Unlock()
			return id, nil
		}

		if buffer.NextReady {
			buffer.SwitchPos()
			buffer.NextReady = false
		}
		buffer.Lock.Unlock()

	}
}

//updateSegmentFromDB 周期制造空洞
func (s *SegmentImpl) updateSegmentFromDB(ctx context.Context, key string, segment *entity.Segment) (err error) {
	var (
		leafAlloc *entity.LeafAlloc
	)
	buffer := segment.Buffer
	if !buffer.IsInitOk() { //第一次初始化
		leafAlloc, err = s.repo.UpdateMaxIdAndGetLeafAlloc(ctx, key)
		if err != nil {
			log.Error("err %+v", err)
			return
		}
		buffer.Step = leafAlloc.Step
		buffer.MinStep = leafAlloc.Step
	} else if buffer.UpdateStamp == 0 { //没有更新过
		leafAlloc, err = s.repo.UpdateMaxIdAndGetLeafAlloc(ctx, key)
		if err != nil {
			log.Error("err %+v", err)
			return
		}
		buffer.UpdateStamp = util.CurrentTimeMillis()
		buffer.Step = leafAlloc.Step
		buffer.MinStep = leafAlloc.Step
	} else { //已经初始化，并且更新过了
		duration := time.Duration(util.CurrentTimeMillis()-buffer.UpdateStamp) * time.Millisecond
		nextStep := buffer.Step
		if duration < SEGMENT_DURATION { //一个周期内
			if nextStep*2 > MAX_STEP {
				//do nothing
			} else {
				nextStep = nextStep * 2
			}
		} else if duration < SEGMENT_DURATION*2 { //两个周期内
			//do nothing with step
		} else { // 两个周期外
			if nextStep/2 >= buffer.MinStep {
				nextStep = nextStep / 2
			}
		}

		tmp := &entity.LeafAlloc{
			BizTag: key,
			Step:   nextStep,
		}
		leafAlloc, err = s.repo.UpdateMaxIdByCustomStepAndGetLeafAlloc(ctx, tmp)
		if err != nil {
			return
		}
		buffer.UpdateStamp = util.CurrentTimeMillis()
		buffer.Step = nextStep
		buffer.MinStep = leafAlloc.Step
	}

	value := leafAlloc.MaxID - buffer.Step
	segment.SetValue(value)
	segment.SetMax(leafAlloc.MaxID)
	segment.SetStep(leafAlloc.Step)

	return nil
}

func (s *SegmentImpl) waitAndSleep(buffer *entity.SegmentBuffer) {
	roll := 0
	for buffer.ThreadRunning.Load() {
		roll += 1
		if roll > 10000 {
			time.Sleep(10 * time.Millisecond)
			break
		}
	}
}
