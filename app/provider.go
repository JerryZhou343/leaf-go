package app

import (
	"github.com/google/wire"
	"github.com/jerryzhou343/leaf-go/app/executor/cmd"
	"github.com/jerryzhou343/leaf-go/app/executor/query"
	"github.com/jerryzhou343/leaf-go/infra/repo"
)

var (
	AppProvider = wire.NewSet(cmd.NewAppCmd, query.NewAppQuery, repo.Provider)
)
