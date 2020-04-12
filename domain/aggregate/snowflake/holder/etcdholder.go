package holder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/jerryzhou343/leaf-go/domain/aggregate/snowflake/valueobj"
	"github.com/jerryzhou343/leaf-go/domain/util"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"time"
)

const (
	requestTimeout = 5 * time.Second
)

type EtcdHolder struct {
	PREFIX_ETCD_PATH string
	PATH_FOREVER     string
	listenAddr       string
	lastUpdateTime   int64
	ip               string
	port             int
	etcdCli          *clientv3.Client
}

func NewEtcdHolder(ip string, port int, srvName string, workerId uint32) (Holder, error) {
	var (
		err error
	)
	ret := &EtcdHolder{}
	ret.port = port
	ret.ip = ip
	ret.listenAddr = ret.ip + fmt.Sprintf(":%d", ret.port)
	ret.PREFIX_ETCD_PATH = "/snowflake/" + srvName
	ret.PATH_FOREVER = ret.PREFIX_ETCD_PATH + "/forever/" + fmt.Sprintf("%d", workerId)

	return ret, err
}

func (e *EtcdHolder) Init(addrs []string) (err error) {
	var (
		rsp *clientv3.GetResponse
	)
	e.etcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:            addrs,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    1 * time.Second,
		DialKeepAliveTimeout: 6 * time.Second,
	})

	if err != nil {
		err = errors.WithStack(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	rsp, err = e.etcdCli.Get(ctx, e.PATH_FOREVER)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Error("cancel by another routine %+v", err)
		case context.DeadlineExceeded:
			log.Error("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
		case rpctypes.ErrKeyNotFound:
			err = nil
		default:
			fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
	} else {
		for _, itr := range rsp.Kvs {
			err = e.checkInitTimeStamp(itr.Value)
		}
	}
	if err == nil {
		go e.ScheduledUploadData()
	}

	fmt.Printf("=>%+v", err)
	return err
}

func (e *EtcdHolder) ScheduledUploadData() {
	for {
		select {
		case <-time.After(5 * time.Second):
			var (
				err error
			)
			if util.CurrentTimeMillis() < e.lastUpdateTime {
				continue
			}
			ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
			_, err = e.etcdCli.Put(ctx, e.PATH_FOREVER, e.buildData())
			cancel()
			if err != nil {
				log.Error("etcd put vale to key %s failed", e.PATH_FOREVER)
				continue
			}
			e.lastUpdateTime = util.CurrentTimeMillis()
		}
	}
}

func (e *EtcdHolder) checkInitTimeStamp(data []byte) (err error) {
	var (
		endpoint valueobj.Endpoint
	)
	json.Unmarshal(data, &endpoint)
	if endpoint.Timestamp > util.CurrentTimeMillis() {
		err = errors.New("node time less tran last timestamp")
	}
	return
}

func (e *EtcdHolder) buildData() string {
	v := valueobj.Endpoint{
		IP:        e.ip,
		Port:      e.port,
		Timestamp: util.CurrentTimeMillis(),
	}
	data, _ := json.Marshal(v)
	return string(data)
}
