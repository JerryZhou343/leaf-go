package entity

import (
	"go.uber.org/atomic"
	"sync"
)

type SegmentBuffer struct {
	Lock          *sync.RWMutex
	currentPos    int        //当前使用的segment的index
	Segments      []*Segment //两个segment
	Key           string
	NextReady     bool
	InitOk        *atomic.Bool
	ThreadRunning *atomic.Bool

	Step        int64
	MinStep     int64
	UpdateStamp int64
}

func NewSegmentBuffer(key string) *SegmentBuffer {
	ret := &SegmentBuffer{
		ThreadRunning: atomic.NewBool(false),
	}
	ret.currentPos = 0
	ret.NextReady = false
	ret.Lock = &sync.RWMutex{}
	ret.Segments = []*Segment{NewSegment(ret), NewSegment(ret)}
	ret.Key = key
	ret.InitOk = atomic.NewBool(false)
	return ret
}

func (s *SegmentBuffer) GetCurrent() *Segment {
	return s.Segments[s.currentPos]
}

func (s *SegmentBuffer) NextPos() int {
	return (s.currentPos + 1) % 2
}

func (s *SegmentBuffer) SwitchPos() {
	s.currentPos = s.NextPos()
}

func (s *SegmentBuffer) IsInitOk() bool {
	return s.InitOk.Load()
}

func (s *SegmentBuffer) SetInitOK(value bool) {
	s.InitOk.Store(value)
}
