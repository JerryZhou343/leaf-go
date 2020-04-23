package app

import (
	"github.com/JerryZhou343/leaf-go/app/executor/cmd"
	"github.com/JerryZhou343/leaf-go/app/executor/query"
	"github.com/JerryZhou343/leaf-go/domain"
	"github.com/JerryZhou343/leaf-go/infra/repo"
	"github.com/google/wire"
)

var (
	AppProvider = wire.NewSet(cmd.NewAppCmd, query.NewAppQuery,
		domain.Provider,
		repo.Provider)
)
