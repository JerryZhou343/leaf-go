package entity

import "go.uber.org/atomic"

type Segment struct {
	Value  *atomic.Int64
	Max    *atomic.Int64
	Step   *atomic.Int64
	Buffer *SegmentBuffer
}

func (s *Segment) GetIdle() int64 {
	return s.Max.Load() - s.Value.Load()
}

func NewSegment(buffer *SegmentBuffer) *Segment {
	return &Segment{
		Value:  atomic.NewInt64(0),
		Max:    atomic.NewInt64(0),
		Step:   atomic.NewInt64(0),
		Buffer: buffer,
	}
}

func (s *Segment) SetValue(value int64) {
	s.Value.Store(value)
}

func (s *Segment) SetMax(value int64) {
	s.Max.Store(value)
}

func (s *Segment) SetStep(value int64) {
	s.Step.Store(value)
}

func (s *Segment) Incr() int64 {
	return s.Value.Add(1)
}

func (s *Segment) GetMax() int64 {
	return s.Max.Load()
}

func (s *Segment) GetStep() int64 {
	return s.Step.Load()
}

func (s *Segment) GetValue() int64 {
	return s.Value.Load()
}
