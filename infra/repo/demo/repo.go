package demo

import (
	"github.com/jerryzhou343/leaf-go/domain/aggregate/segment"
	"github.com/jerryzhou343/leaf-go/infra/conf"
)

type repo struct {
}

func NewRepo(conf *conf.Config) (segment.Repo, error) {
	return &repo{},nil
}