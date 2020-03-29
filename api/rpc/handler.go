package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jerryzhou343/leaf-go/app/executor/cmd"
	"github.com/jerryzhou343/leaf-go/app/executor/query"
	"github.com/jerryzhou343/leaf-go/genproto/github.com/jerryzhou343/leaf-go/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	cmdApp   *cmd.AppCmd
	queryApp *query.AppQuery
}

func NewHandler(cmdApp *cmd.AppCmd, queryApp *query.AppQuery) *Handler {
	return &Handler{
		cmdApp:   cmdApp,
		queryApp: queryApp,
	}
}

func (h *Handler) Ping(ctx context.Context, req *empty.Empty) (rsp *empty.Empty, err error) {
	rsp = &empty.Empty{}
	err = status.Error(codes.OK, "ok")
	return
}
