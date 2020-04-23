package query

import (
	"github.com/JerryZhou343/leaf-go/domain/aggregate/segment"
)

type AppQuery struct {
	demoRepo segment.Repo
}

func NewAppQuery(demoRepo segment.Repo) *AppQuery {
	return &AppQuery{
		demoRepo: demoRepo,
	}
}
