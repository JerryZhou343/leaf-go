package cmd

import (
	"github.com/jerryzhou343/leaf-go/domain/aggregate/snowflake"
	"github.com/jerryzhou343/leaf-go/infra/conf"
)

type AppCmd struct {
	//demoRepo segment.Repo
	snowflakeIdGen *snowflake.Snowflake
}

func NewAppCmd(config *conf.Config /*demoRepo segment.Repo,*/, snowflakeIdGen *snowflake.Snowflake) (ret *AppCmd, err error) {

	ret = &AppCmd{
		snowflakeIdGen: snowflakeIdGen,
	}
	return ret, err
}

func (a *AppCmd) GetSnowflakeID(key string) (id int64, err error) {
	return a.snowflakeIdGen.Get(key)
}
