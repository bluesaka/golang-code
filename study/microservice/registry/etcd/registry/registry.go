package registry

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type ServiceInfo struct {
	Name string
	IP   string
}

type Service struct {
	ServiceInfo ServiceInfo
	stop        chan error
	leaseID     clientv3.LeaseID
	client      *clientv3.Client
}

// NewService 创建一个注册服务
func NewService(serviceInfo ServiceInfo, endpoints []string) (service *Service, err error) {
	// clientv3.New() won't return error when no endpoint is available
	// 在v3.3+版本使用了无效的etcd服务短地址时，即使设置了DialTimeout，也不会返回任何错误
	// 可以使用Status查看状态
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          2 * time.Second,
		AutoSyncInterval:     time.Minute,
		DialKeepAliveTime:    5 * time.Second,
		DialKeepAliveTimeout: 5 * time.Second,
		RejectOldCluster:     true,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// 使用Status来检测endpoint
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = client.Status(ctx, endpoints[0])
	if err != nil {
		log.Printf("endpoint: %s is not available\n", endpoints[0])
		panic(err)
	}

	service = &Service{
		ServiceInfo: serviceInfo,
		client:      client,
	}
	return
}

func (service Service) Start() (err error) {
	ch, err := service.keepAlive()
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		select {
		case err := <-service.stop:
			return err
		case <-service.client.Ctx().Done():
			log.Println("context canceled")
			return errors.New("context canceled")
		case resp, ok := <-ch:
			if !ok {
				log.Println("keep alive channel closed")
				return service.revoke()
			}
			log.Printf("recv reply from service: %s, ttl: %d\n", service.getKey(), resp.TTL)
		}
	}
}

// keepAlive
func (service *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	info := service.ServiceInfo
	key := info.Name + "/" + info.IP
	//key := info.Name
	val, _ := jsoniter.Marshal(info)
	ctx := context.Background()

	resp, err := service.client.Grant(ctx, 5)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = service.client.Put(ctx, key, string(val), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	service.leaseID = resp.ID
	return service.client.KeepAlive(ctx, resp.ID)
}

func (service *Service) Stop() {
	service.stop <- nil
}

func (service *Service) getKey() string {
	return service.ServiceInfo.Name // + "/" + service.ServiceInfo.IP
}

func (service *Service) revoke() error {
	_, err := service.client.Revoke(context.Background(), service.leaseID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("service: %s stop\n", service.getKey())
	return err
}
