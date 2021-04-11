package factory

import (
	"errors"
	"sync"
)

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, error)
}

type RedisCache struct {
	data map[string]interface{}
	mux sync.Mutex
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if v, ok := r.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("redis no value")
}

func (r *RedisCache) Set(key string, value interface{}) bool {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.data[key] = value
	return true
}

type Memcache struct {
	data map[string]interface{}
	mux sync.Mutex
}

func (r *Memcache) Get(key string) (interface{}, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if v, ok := r.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("memcache no value")
}

func (r *Memcache) Set(key string, value interface{}) bool {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.data[key] = value
	return true
}

type CacheType int

const (
	redis CacheType = iota
	memcache
)

type CacheFactory struct {}

func (c *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	switch cacheType {
	case redis:
		return &RedisCache{data: map[string]interface{}{}}, nil
	case memcache:
		return &Memcache{data: map[string]interface{}{}}, nil
	}
	return nil, errors.New("not support cache type")
}

