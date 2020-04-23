package domain

import (
	"github.com/JerryZhou343/leaf-go/domain/aggregate/segment"
	"github.com/JerryZhou343/leaf-go/domain/aggregate/snowflake"
	"github.com/google/wire"
)

var (
	Provider = wire.NewSet(
		snowflake.NewSnowflake,
		segment.NewSegmentImpl,
	)
)
