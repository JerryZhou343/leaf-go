package mysql

import (
	"github.com/JerryZhou343/leaf-go/infra/conf"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
)

func NewMySQL(c *conf.Config) (db *sql.DB, err error) {
	if c.SQL.QueryTimeout == 0 || c.SQL.ExecTimeout == 0 || c.SQL.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err = sql.Open(c.SQL)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}
