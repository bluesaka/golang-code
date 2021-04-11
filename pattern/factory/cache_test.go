package factory

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	cacheFactory := &CacheFactory{}
	fmt.Println(redis)
	redis, err := cacheFactory.Create(redis)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(redis)
	redis.Set("k1", "v1")
	fmt.Println(redis.Get("k1"))
}
