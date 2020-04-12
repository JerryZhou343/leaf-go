package test

import (
	"context"
	"fmt"
	leaf_go "github.com/jerryzhou343/leaf-go/genproto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func BenchmarkSegment(b *testing.B) {

	cc, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		b.Errorf("dial failed 【%+v]", err)
		return
	}
	client := leaf_go.NewLeafGoServiceClient(cc)
	req := &leaf_go.GetSegmentReq{
		Key: "leaf-segment-test",
	}
	//var 	rsp  *leaf_go.GetSegmentRsp
	defer func(begin time.Time) {
		fmt.Printf("time cost [%+v]", time.Now().Sub(begin).Seconds())
	}(time.Now())
	for i := 0; i < 100000; i++ {
		_, err = client.GetSegment(context.Background(), req)
		if err != nil {
			b.Errorf("err %+v", err)
		}
	}

}

func TestSegment(t *testing.T) {
	var (
		set map[int64]struct{}
	)
	set = map[int64]struct{}{}
	cc, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		t.Errorf("dial failed 【%+v]", err)
		return
	}
	client := leaf_go.NewLeafGoServiceClient(cc)
	req := &leaf_go.GetSegmentReq{
		Key: "leaf-segment-test",
	}
	var rsp *leaf_go.GetSegmentRsp
	defer func(begin time.Time) {
		fmt.Printf("time cost [%+v]", time.Now().Sub(begin).Seconds())
	}(time.Now())
	for i := 0; i < 100000; i++ {
		rsp, err = client.GetSegment(context.Background(), req)
		if err != nil {
			t.Errorf("err %+v", err)
			break
		}
		if _, ok := set[rsp.Id]; ok {
			t.Errorf("duplicate id %v", rsp.Id)
			break
		} else {
			set[rsp.Id] = struct{}{}
		}
	}

}
