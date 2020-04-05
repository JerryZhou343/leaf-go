// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/jerryzhou343/leaf-go/api"
	"github.com/jerryzhou343/leaf-go/infra/conf"
	"github.com/jerryzhou343/leaf-go/infra/server/grpc"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(conf.NewConf,api.Provider, grpc.New, NewApp))
}
