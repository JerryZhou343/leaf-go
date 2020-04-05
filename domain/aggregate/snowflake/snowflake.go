package snowflake

import (
	"github.com/bilibili/kratos/pkg/net/ip"
	hold "github.com/jerryzhou343/leaf-go/domain/aggregate/snowflake/holder"
	"github.com/jerryzhou343/leaf-go/domain/aggregate/snowflake/idgen"
	"github.com/jerryzhou343/leaf-go/infra/conf"
	"strconv"
	"strings"
)

type Snowflake struct {
	generator *idgen.SnowflakeIDGen
	holder    hold.Holder
	config    *conf.Config
}

func NewSnowflake(config *conf.Config) (ret *Snowflake, err error) {
	ret = &Snowflake{}
	ret.generator, err = idgen.NewSnowflakeIDGenImpl(int64(config.Snowflake.WorkerId))
	addrInfo := strings.Split(config.Grpc.Addr, ":")
	port, err := strconv.Atoi(addrInfo[1])
	ret.holder, err = hold.NewEtcdHolder(ip.InternalIP(), port, config.Snowflake.SrvName, config.Snowflake.WorkerId)
	err = ret.holder.Init(config.Snowflake.EtcdCluster)
	return ret, err
}

func (s *Snowflake) Get(key string) (int64, error) {
	return s.generator.Get(key)
}
