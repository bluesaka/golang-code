/**
@link https://www.zhihu.com/zvideo/1308438116354207744

- G: 一个goroutine
- P: processor处理器，默认gomaxprocs个，有多少个P，决定能同时并发执行的goroutine
- M: os线程，真正执行代码的线程


程序启动时，先初始化和核数一样的P。
当创建一个goroutine时，优先尝试放在本地队列，如果本地队列满了，则会把本地队列的前半部分和这个新的goroutine一起移到全局队列中。
如果没有可用的P的时候，新goroutine加入全局队列中。
如果获取到空闲的P，那么尝试去唤醒一个M，没有可用的M的时候新建一个M。
当M关联上P时，且local队列有任务时，可以一直从p的local队列中取goroutine执行。
当P的local队列中没有goroutine时，则会尝试从全局队列中拿一部分放在本地队列中，这个过程是加锁的。
当从全局队列没取到时，会尝试从其他的P的local队列偷取一半放在自己的本地队列中
当一个G发生系统调用的时候，P会断开与当前的M的关系，尝试从M的空闲队列中获取一个M来继续执行剩下的goroutine。
当上面的G系统调用结束后，M尝试获取一个P来继续执行，如果没获取到，则会把这个g放到全局队列中，并且自己进入M的空闲队列。这里不是销毁M，避免后面又要创建M，造成不必要的开销。

 */

package gpm