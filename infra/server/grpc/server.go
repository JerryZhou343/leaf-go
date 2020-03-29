package grpc

import (
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/jerryzhou343/leaf-go/infra/conf"
	stub "github.com/jerryzhou343/leaf-go/genproto/github.com/jerryzhou343/leaf-go/v1"
)

// New new a grpc server.
func New(svc stub.LeafGoServiceServer, config *conf.Config) (ws *warden.Server, err error) {
	ws = warden.NewServer(config.Grpc)
	stub.RegisterLeafGoServiceServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
