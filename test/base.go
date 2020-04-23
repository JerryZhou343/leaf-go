package test

import (
	"context"
	"fmt"
	leaf_go "github.com/JerryZhou343/leaf-go/genproto"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

// AppID .
const AppID = "leaf-go"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (leaf_go.LeafGoServiceClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return leaf_go.NewLeafGoServiceClient(cc), nil
}
