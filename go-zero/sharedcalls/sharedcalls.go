/**
@link https://hub.fastgit.org/tal-tech/zero-doc/blob/main/doc/sharedcalls.md

并发场景下，可能会有多个线程（协程）同时请求同一份资源，如果每个请求都要走一遍资源的请求过程，除了比较低效之外，还会对资源服务造成并发的压力。
举一个具体例子，比如缓存失效，多个请求同时到达某服务请求某资源，该资源在缓存中已经失效，此时这些请求会继续访问DB做查询，会引起数据库压力瞬间增大。
而使用SharedCalls可以使得同时多个请求只需要发起一次拿结果的调用，其他请求"坐享其成"，这种设计有效减少了资源服务的并发压力，可以有效防止缓存击穿。
*/
package sharedcalls

import "sync"

type MySharedCalls interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
}

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type mySharedGroup struct {
	calls map[string]*call
	lock  sync.Mutex
}

func NewMySharedCall() MySharedCalls {
	return &mySharedGroup{
		calls: map[string]*call{},
	}
}

func (g *mySharedGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.lock.Lock()
	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := g.makeCall(key, fn)
	return c.val, c.err
}

func (g *mySharedGroup) makeCall(key string, fn func() (interface{}, error)) *call {
	c := new(call)
	c.wg.Add(1)
	g.calls[key] = c
	g.lock.Unlock()

	defer func() {
		g.lock.Lock()
		delete(g.calls, key)
		g.lock.Unlock()

		c.wg.Done()
	}()

	c.val, c.err = fn()
	return c
}
