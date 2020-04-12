package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jerryzhou343/leaf-go/app/executor/cmd"
	"github.com/jerryzhou343/leaf-go/app/executor/query"
	stub "github.com/jerryzhou343/leaf-go/genproto"
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

func (h *Handler) GetSnowflake(ctx context.Context, req *stub.GetSnowflakeReq) (rsp *stub.GetSnowflakeRsp, err error) {
	rsp = &stub.GetSnowflakeRsp{}
	rsp.Id, err = h.cmdApp.GetSnowflakeID(req.Key)
	return
}

func (h *Handler) GetSegment(ctx context.Context, req *stub.GetSegmentReq) (rsp *stub.GetSegmentRsp, err error) {
	rsp = &stub.GetSegmentRsp{}
	rsp.Id, err = h.cmdApp.GetSegmentID(ctx, req.Key)
	return
}
