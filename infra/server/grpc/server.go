package grpc

import (
	stub "github.com/JerryZhou343/leaf-go/genproto"
	"github.com/JerryZhou343/leaf-go/infra/conf"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc stub.LeafGoServiceServer, config *conf.Config) (ws *warden.Server, err error) {
	ws = warden.NewServer(config.Grpc)
	stub.RegisterLeafGoServiceServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
