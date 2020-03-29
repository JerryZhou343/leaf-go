package cmd

import (
	"github.com/jerryzhou343/leaf-go/domain/aggregate/segment"
)

type AppCmd struct {
	demoRepo segment.Repo
}

func NewAppCmd(demoRepo segment.Repo) *AppCmd {
	return &AppCmd{
		demoRepo: demoRepo,
	}
}
