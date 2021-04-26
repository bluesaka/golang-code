/**
断点调试汇编
目前Go语言支持GDB、LLDB和Delve几种调试器。只有Delve是专门为Go语言设计开发的调试工具，而且Delve本身也是采用Go语言开发
go get github.com/go-delve/delve/cmd/dlv
会安装在$GOPATH/bin下，生成dlv可执行文件

> dlv debug
> break main.main
> breakpoints
> continue
> disassemble
> break runtime.newobject
> print typ
> args
> locals

(dlv) break runtime.newobject
Breakpoint 2 set at 0x100dcef for runtime.newobject() /usr/local/Cellar/go/1.15.5/libexec/src/runtime/malloc.go:1194
(dlv) continue
> runtime.newobject() /usr/local/Cellar/go/1.15.5/libexec/src/runtime/malloc.go:1194 (hits goroutine(1):1 total:1) (PC: 0x100dcef)
Warning: debugging optimized function
  1189:	}
  1190:
  1191:	// implementation of new builtin
  1192:	// compiler (both frontend and SSA backend) knows the signature
  1193:	// of this function
=>1194:	func newobject(typ *_type) unsafe.Pointer {
  1195:		return mallocgc(typ.size, typ, true)
  1196:	}
  1197:
  1198:	//go:linkname reflect_unsafe_New reflect.unsafe_New
  1199:	func reflect_unsafe_New(typ *_type) unsafe.Pointer {
(dlv) print typ
*runtime._type {size: 16, ptrdata: 8, hash: 875453117, tflag: tflagUncommon|tflagExtraStar|tflagNamed (7), align: 8, fieldAlign: 8, kind: 25, equal: runtime.strequal, gcdata: *1, str: 4950, ptrToThis: 28384}
可以看到这里打印的size是16bytes，因为A结构体里只有一个string类型的字段
*/
package main

import "fmt"

type A struct {
	name string
}

func main() {
	a := new(A)
	fmt.Println(a)
}
