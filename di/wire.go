// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/JerryZhou343/leaf-go/api"
	"github.com/JerryZhou343/leaf-go/infra/conf"
	"github.com/JerryZhou343/leaf-go/infra/server/grpc"
	"github.com/google/wire"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(conf.NewConf, api.Provider, grpc.New, NewApp))
}
