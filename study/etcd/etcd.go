package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func Etcd1() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.Put(ctx, "test-key", "value1")
	if err != nil {
		log.Println("etcd put failed, err:", err)
	}

	resp, err := client.Get(ctx, "test-key")
	if err != nil {
		log.Println("etcd get failed, err:", err)
	}
	log.Println(resp)
	for _, v := range resp.Kvs {
		log.Printf("key: %s, value: %s, version: %d\n", v.Key, v.Value, v.Version)
	}
}

func Watch() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// clientv3.WithXXX
	watchChan := client.Watch(context.Background(), "test-key")
	log.Println("etcd start watching")

	for resp := range watchChan {
		for _, v := range resp.Events {
			log.Printf("watch type: %s, key: %s, value: %s", v.Type, v.Kv.Key, v.Kv.Value)
		}
	}
}

func Lease() {

}
