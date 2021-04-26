package abstractfactory

import (
	"errors"
	"sync"
)

type Cache2 interface {
	Set2(key string, value interface{}) bool
	Get2(key string) (interface{}, error)
}

type RedisCache2 struct {
	data map[string]interface{}
	mux  sync.Mutex
}

func (r *RedisCache2) Get2(key string) (interface{}, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if v, ok := r.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("redis no value")
}

func (r *RedisCache2) Set2(key string, value interface{}) bool {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.data[key] = value
	return true
}

type Memcache2 struct {
	data map[string]interface{}
	mux  sync.Mutex
}

func (r *Memcache2) Get2(key string) (interface{}, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if v, ok := r.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("memcache no value")
}

func (r *Memcache2) Set2(key string, value interface{}) bool {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.data[key] = value
	return true
}

// 抽象工厂
type CacheAbstractFactory interface {
	Create() Cache2
}

type RedisFactory struct {
}

func (r *RedisFactory) Create() Cache2 {
	return &RedisCache2{data: map[string]interface{}{}}
}

type MemcacheFactory struct {
}

func (r *MemcacheFactory) Create() Cache2 {
	return &Memcache2{data: map[string]interface{}{}}
}
