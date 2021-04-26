/**
@link https://www.liwenzhou.com/posts/Go/performance_optimisation/
Go性能优化

Go性能优化主要有以下几个方面：
- CPU profile：报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据
- Memory Profile（Heap Profile）：报告程序的内存使用情况
- Block Profiling：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈
- Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的

go tool pprof cpu.pprof

"runtime/pprof"
_ "net/http/pprof"

go-torch
github.com/uber/go-torch

graphviz图形化工具
brew install graphviz

火焰图 (Flame Graph)
*/

package pprof

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var CPUPprofSwitch = flag.Bool("cpu", false, "cpu pprof switch")
var MemPprofSwitch = flag.Bool("mem", false, "mem pprof switch")

func Pprof1() {
	flag.Parse()

	if *CPUPprofSwitch {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			log.Println("create cpu pprof failed, err:", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 8; i++ {
		go func1()
	}

	time.Sleep(time.Second * 20)

	if *MemPprofSwitch {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			log.Println("create mem pprof failed, err:", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

func func1() {
	var c chan int
	for {
		select {
		case v := <-c:
			log.Println("recv from chan, value=", v)
		}
	}
}
