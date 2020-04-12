package repo

import (
	"github.com/google/wire"
	"github.com/jerryzhou343/leaf-go/infra/driver/mysql"
	"github.com/jerryzhou343/leaf-go/infra/repo/segment"
)

var (
	Provider = wire.NewSet(mysql.NewMySQL, segment.NewRepo)
)
