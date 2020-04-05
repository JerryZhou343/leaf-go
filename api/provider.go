package api

import (
	"github.com/google/wire"
	"github.com/jerryzhou343/leaf-go/api/rpc"
	"github.com/jerryzhou343/leaf-go/app"
	stub "github.com/jerryzhou343/leaf-go/genproto"
)

var (
	Provider = wire.NewSet(rpc.NewHandler, wire.Bind(new(stub.LeafGoServiceServer), new(*rpc.Handler)),
		app.AppProvider,
	)
)
