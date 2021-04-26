/**
时间轮 TimingWheel
@link https://mp.weixin.qq.com/s/CiZ5SpuT-VN8V9wil8_iGg
@link https://go-zero.dev/timing-wheel.html
*/

package timingwheel

import (
	"github.com/tal-tech/go-zero/core/collection"
	"log"
	"runtime"
	"time"
)

func Cache1() {
	log.Println("num goroutine", runtime.NumGoroutine())
	// cache的timingWheel，控制内存缓存一秒后过期
	cache, err := collection.NewCache(time.Second, collection.WithName("in-memory cache"))
	if err != nil {
		log.Println("NewCache err:", err)
	}
	// 多出两个协程，一个协程处理缓存的统计，一个协程处理时间轮任务
	log.Println("num goroutine", runtime.NumGoroutine())

	cache.Set("a", "a")
	log.Println(cache.Get("a"))
	time.Sleep(time.Second)
	log.Println(cache.Get("a"))

	// expired
	time.Sleep(time.Millisecond)
	log.Println(cache.Get("a"))
}
