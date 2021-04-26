/**
Go的调度器 GPM (Goroutine, Processor, Machine)
M是操作系统线程，在绝大多数情况下，P和M的数量相等，每创建一个P，就会创建一个M，只有少数情况下M的数量会大于P

runtime.GOMAXPROCS(n int) 来设置P的值
Go1.5开始，GOMAXPROCS默认值为CPU的核数
但对于IO密集型的场景，把GOMAXPROCS的值适当调大，比如两倍的CPU核数，提高并发的运行性能
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpu count", runtime.NumCPU())
	fmt.Println("original GOMAXPROCS", runtime.GOMAXPROCS(-1))
	fmt.Println("cpu GOMAXPROCS", runtime.GOMAXPROCS(runtime.NumCPU()))

	runtime.GOMAXPROCS(10)
	fmt.Println("after GOMAXPROCS", runtime.GOMAXPROCS(-1))
}
