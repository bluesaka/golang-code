package abstractfactory

import (
	"fmt"
	"testing"
)

func Test_Abstract_Cache_Create(t *testing.T) {
	redisFactory := &RedisFactory{}
	fmt.Println(redisFactory)
	redis := redisFactory.Create()
	redis.Set2("k2", "v2")
	fmt.Println(redis.Get2("k2"))
}
