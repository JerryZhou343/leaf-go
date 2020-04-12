package domain

import (
	"github.com/google/wire"
	"github.com/jerryzhou343/leaf-go/domain/aggregate/segment"
	"github.com/jerryzhou343/leaf-go/domain/aggregate/snowflake"
)

var (
	Provider = wire.NewSet(
		snowflake.NewSnowflake,
		segment.NewSegmentImpl,
	)
)
