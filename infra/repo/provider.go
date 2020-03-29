package repo

import (
	"github.com/google/wire"
	"github.com/jerryzhou343/leaf-go/infra/repo/demo"
)

var (
	Provider = wire.NewSet(demo.NewRepo)
)
