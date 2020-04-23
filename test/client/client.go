package main

import (
	"context"
	"fmt"
	leaf_go "github.com/JerryZhou343/leaf-go/genproto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var (
		key     string
		rootCmd = &cobra.Command{}
		segment = &cobra.Command{
			Use: "segment",
			Run: func(cmd *cobra.Command, args []string) {
				if key != "" {
					TestSegment(key)
				}
			},
		}
	)
	rootCmd.AddCommand(segment)
	segment.PersistentFlags().StringVarP(&key, "key", "", "", "segment key")
	rootCmd.Execute()
}

func TestSegment(key string) {
	var (
		set map[int64]struct{}
	)
	set = map[int64]struct{}{}
	cc, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("dial failed 【%+v]", err)
		return
	}
	client := leaf_go.NewLeafGoServiceClient(cc)
	req := &leaf_go.GetSegmentReq{
		Key: key,
	}
	var rsp *leaf_go.GetSegmentRsp
	defer func(begin time.Time) {
		fmt.Printf("time cost [%+v]", time.Now().Sub(begin).Seconds())
	}(time.Now())
	for i := 0; i < 100000; i++ {
		rsp, err = client.GetSegment(context.Background(), req)
		if err != nil {
			log.Printf("err %+v", err)
			break
		}
		if _, ok := set[rsp.Id]; ok {
			log.Printf("duplicate id %v", rsp.Id)
			break
		} else {
			set[rsp.Id] = struct{}{}
		}
	}

}
