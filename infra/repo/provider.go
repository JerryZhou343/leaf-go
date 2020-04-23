package repo

import (
	"github.com/JerryZhou343/leaf-go/infra/driver/mysql"
	"github.com/JerryZhou343/leaf-go/infra/repo/segment"
	"github.com/google/wire"
)

var (
	Provider = wire.NewSet(mysql.NewMySQL, segment.NewRepo)
)
