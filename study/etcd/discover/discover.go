package discover

import (
	"context"
	"fmt"
	mvccpb2 "github.com/coreos/etcd/mvcc/mvccpb"
	jsoniter "github.com/json-iterator/go"
	"go-code/study/etcd/registry"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/resolver"
	"log"
)

const schema = "etcd"

// Resolver is the implementation of grpc.resolve.Builder
// Resolver 实现grpc的grpc.resolve.Builder接口的Build与Scheme方法
type Resolver struct {
	endpoints []string
	service   string
	cli       *clientv3.Client
	cc        resolver.ClientConn
}

// NewResolver return resolver builder
// endpoints example: http://127.0.0.1:2379 http://127.0.0.1:12379 http://127.0.0.1:22379"
// service is service name
func NewResolver(endpoints []string, service string) resolver.Builder {
	return &Resolver{endpoints: endpoints, service: service}
}

func (r *Resolver) watch(prefix string) {
	addrs := make(map[string]resolver.Address)

	update := func() {
		addrList := make([]resolver.Address, 0, len(addrs))
		for _, v := range addrs {
			addrList = append(addrList, v)
		}
		r.cc.UpdateState(resolver.State{Addresses: addrList})
	}

	resp, err := r.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err == nil {
		for i, kv := range resp.Kvs {
			var info registry.ServiceInfo
			err := jsoniter.Unmarshal(kv.Value, &info)
			if err != nil {
				log.Println("kvs unmarshal error", err)
			} else {
				addrs[string(resp.Kvs[i].Value)] = resolver.Address{Addr: info.IP}
			}
		}
	}

	update()

	watchChan := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range watchChan {
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb2.Event_EventType(mvccpb.PUT):
				info := &registry.ServiceInfo{}
				err := jsoniter.Unmarshal(ev.Kv.Value, info)
				if err != nil {
					log.Println("ev json unmarshal error", err)
				} else {
					addrs[string(ev.Kv.Key)] = resolver.Address{Addr: info.IP}
				}
			case mvccpb2.Event_EventType(mvccpb.DELETE):
				delete(addrs, string(ev.PrevKv.Key))
			}
		}
		update()
	}
}

func (r Resolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (r Resolver) Close() {
}

func (r Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: r.endpoints,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc: create clientv3 client failed: %v", err)
	}

	r.cli = cli
	r.cc = cc

	return r, nil
}

func (r Resolver) Scheme() string {
	return schema + "_" + r.service
}
