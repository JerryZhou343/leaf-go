package conf

import (
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// Config .
type Config struct {
	Redis     *redis.Config
	SQL       *sql.Config
	Grpc      *warden.ServerConfig
	App       *ApplicationConfig
	Snowflake *SnowflakeConfig
}

type ApplicationConfig struct {
}

type SnowflakeConfig struct {
	EtcdCluster []string
	WorkerId    uint32
	SrvName     string
}

func NewConf() (conf *Config, err error) {
	var (
		sqlct         paladin.TOML
		rpcCt         paladin.TOML
		redisCt       paladin.TOML
		applicationCt paladin.TOML
	)

	conf = &Config{
		Redis:     &redis.Config{},
		SQL:       &sql.Config{},
		Grpc:      &warden.ServerConfig{},
		App:       &ApplicationConfig{},
		Snowflake: &SnowflakeConfig{},
	}
	//rpc server
	if err = paladin.Get("grpc.toml").Unmarshal(&rpcCt); err != nil {
		return
	}

	if err = rpcCt.Get("Server").UnmarshalTOML(conf.Grpc); err != nil {
		return
	}

	//sql
	if err = paladin.Get("db.toml").Unmarshal(&sqlct); err != nil {
		return
	}
	if err = sqlct.Get("Client").UnmarshalTOML(conf.SQL); err != nil {
		return
	}

	//redis
	if err = paladin.Get("redis.toml").Unmarshal(&redisCt); err != nil {
		return
	}

	if err = redisCt.Get("Client").UnmarshalTOML(&conf.Redis); err != nil {
		return
	}
	//application
	if err = paladin.Get("application.toml").Unmarshal(&applicationCt); err != nil {
		return
	}

	if err = applicationCt.Get("Application").UnmarshalTOML(&conf.App); err != nil {
		return
	}
	//snowflake
	if err = applicationCt.Get("Snowflake").UnmarshalTOML(&conf.Snowflake); err != nil {
		return
	}
	return
}
