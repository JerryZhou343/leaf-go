package cmd

import (
	"context"
	"github.com/jerryzhou343/leaf-go/domain/aggregate/segment"
	"github.com/jerryzhou343/leaf-go/domain/aggregate/snowflake"
	"github.com/jerryzhou343/leaf-go/infra/conf"
)

type AppCmd struct {
	snowflakeIdGen *snowflake.Snowflake
	segmentImpl    *segment.SegmentImpl
	config         *conf.Config
}

func NewAppCmd(config *conf.Config, segmentImpl *segment.SegmentImpl, snowflakeIdGen *snowflake.Snowflake) (ret *AppCmd, err error) {
	err = segmentImpl.Init()
	if err != nil {
		return
	}
	//todo:snowflake 按照配置条件构造
	ret = &AppCmd{
		snowflakeIdGen: snowflakeIdGen,
		segmentImpl:    segmentImpl,
		config:         config,
	}

	return ret, err
}

func (a *AppCmd) GetSnowflakeID(key string) (id int64, err error) {
	return a.snowflakeIdGen.Get(key)
}

func (a *AppCmd) GetSegmentID(ctx context.Context, key string) (id int64, err error) {
	return a.segmentImpl.Get(ctx, key)
}
