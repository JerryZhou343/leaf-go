package idgen

import (
	"github.com/jerryzhou343/leaf-go/domain/util"
	stub "github.com/jerryzhou343/leaf-go/genproto"
	"math/rand"
	"time"
)

type SnowflakeIDGen struct {
	twepoch            int64      //起止时间戳
	workerId           int64      //工作者ID
	lastTimestamp      int64      //最后刷新时间
	workerIdBits       int64      //worker id 位数
	maxWorkerId        int64      //worker id 最大值
	sequenceBits       int64      //序列号位数
	workerIdShift      int64      //worker id 左移位数
	timestampLeftShift int64      //timestamp 左移位数
	sequenceMask       int64      //sequence 掩码
	sequence           int64      //sequence
	random             *rand.Rand //随机数种子
}

func NewSnowflakeIDGenImpl(workerID int64) (*SnowflakeIDGen, error) {
	var (
		ret *SnowflakeIDGen
		err error
	)
	ret = &SnowflakeIDGen{
		sequenceBits:  12,
		workerIdBits:  10,
		lastTimestamp: -1,
		twepoch:       1288834974657,
	}
	ret.workerIdShift = ret.sequenceBits
	ret.timestampLeftShift = ret.sequenceBits + ret.workerIdBits
	if ret.twepoch > util.CurrentTimeMillis() {
		err = stub.LAST_TIME_GT_CURRENT_TIME
		return nil, err
	}
	ret.sequenceMask = int64(^(-1 << ret.sequenceBits))
	ret.maxWorkerId = ^(-1 << ret.workerId)
	ret.workerId = workerID
	ret.sequence = 0
	ret.random = rand.New(rand.NewSource(time.Now().UnixNano()))

	return ret, nil
}

func (s *SnowflakeIDGen) Get(key string) (int64, error) {
	//当前时间
	timestamp := util.CurrentTimeMillis()
	if timestamp < s.lastTimestamp {
		offset := s.lastTimestamp - timestamp
		//当前时间晚于上一次更新时间
		if offset <= 5 {
			waitDuration := time.Duration(offset<<1) * time.Millisecond
			select {
			case <-time.After(waitDuration):
				timestamp = util.CurrentTimeMillis()
				if timestamp < s.lastTimestamp {
					return 0, stub.LAST_TIME_GT_CURRENT_TIME
				}
			}

		}
	}

	//如果和上一次时间是同一时间
	if s.lastTimestamp == timestamp {
		s.sequence = (s.sequence + 1) & s.sequenceMask
		if s.sequence == 0 {
			//seq 为0的时候表示是下一毫秒时间开始对seq做随机
			s.sequence = s.random.Int63n(100)
			timestamp = s.tilNextMillis(s.lastTimestamp)
		}
	} else {
		//如果是新的ms开始
		s.sequence = s.random.Int63n(100)
	}
	s.lastTimestamp = timestamp
	id := ((timestamp - s.twepoch) << s.timestampLeftShift) | (s.workerId << s.workerIdShift) | s.sequence
	return id, nil
}

func (s *SnowflakeIDGen) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := util.CurrentTimeMillis()
	for timestamp <= s.lastTimestamp {
		timestamp = util.CurrentTimeMillis()
	}
	return timestamp
}
