package api

import (
	"github.com/JerryZhou343/leaf-go/api/rpc"
	"github.com/JerryZhou343/leaf-go/app"
	stub "github.com/JerryZhou343/leaf-go/genproto"
	"github.com/google/wire"
)

var (
	Provider = wire.NewSet(rpc.NewHandler, wire.Bind(new(stub.LeafGoServiceServer), new(*rpc.Handler)),
		app.AppProvider,
	)
)
