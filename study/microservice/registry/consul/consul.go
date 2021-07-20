package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
)

func KV() {
	// Get a new client
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &consulapi.KVPair{Key: "consul-test", Value: []byte("abc123")}
	writeMeta, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("writeMeta: %+v\n", writeMeta)

	// Lookup the pair
	pair, queryMeta, err := kv.Get("consul-test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("kv: %v %s\n", pair.Key, pair.Value)
	fmt.Printf("pair: %+v, queryMeta: %+v\n", pair, queryMeta)
}

func RegisterServer() {
	// Get a new client
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "test_node_1"
	registration.Name = "test_node"
	registration.Port = 8600
	registration.Tags = []string{"node tag"}
	registration.Address = "127.0.0.1"
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           "http://127.0.0.1:9001/check",
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	if err := client.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "check")
	})
	http.ListenAndServe(":9001", nil)
}
