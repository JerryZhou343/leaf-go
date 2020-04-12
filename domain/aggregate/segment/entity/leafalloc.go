package entity

import "time"

type LeafAlloc struct {
	BizTag      string
	MaxID       int64
	Step        int64
	Description string
	UpdateTime  time.Time
}
