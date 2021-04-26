/**
@link https://www.cnblogs.com/luozhiyun/p/14349331.html

golang的内存分配基于TCMalloc
mallocgc在分配内存的时候，会按照对象的大小分为3档来进行分配：

- 小对象(tiny): < 16bytes；
	- 使用mcache的tiny分配器分配
- 普通对象: 16bytes ~ 32K；
	- 计算对象的大小，然后使用mcache中相应规格大小的mspan分配
	- 如果 mcache   中没有相应规格大小的mspan，则向mcentral申请
	- 如果 mcentral 中没有相应规格大小的mspan，则向mheap申请
	- 如果 mheap    中没有相应规格大小的mspan，则向操作系统申请
- 大对象: > 32K；
	- 通过largeAlloc来分配一个mspan，直接向mheap申请，当大对象需要分配的页数小于16页时，会直接从pageCache中分配，否则才会从堆页中获取。

*/
package main

import "fmt"

func main() {
	fmt.Println(*test())
}

func test() *int {
	x := new(int)
	*x = 0xAABB
	return x
}

/**
# felix in ~/gosrc/go-code/study/bigobject/cmd [10:53:16] C:1
$ go build -o main main.go

# felix in ~/gosrc/go-code/study/bigobject/cmd [10:53:35]
$ go tool objdump -s main.main main
TEXT main.main(SB) /Users/felix/gosrc/go-code/study/bigobject/cmd/main.go
  main.go:22		0x10a6e80		65488b0c2530000000	MOVQ GS:0x30, CX
  main.go:22		0x10a6e89		483b6110		CMPQ 0x10(CX), SP
  main.go:22		0x10a6e8d		0f867c000000		JBE 0x10a6f0f
  main.go:22		0x10a6e93		4883ec58		SUBQ $0x58, SP
  main.go:22		0x10a6e97		48896c2450		MOVQ BP, 0x50(SP)
  main.go:22		0x10a6e9c		488d6c2450		LEAQ 0x50(SP), BP
  main.go:23		0x10a6ea1		48c70424bbaa0000	MOVQ $0xaabb, 0(SP)
  main.go:23		0x10a6ea9		e8322ff6ff		CALL runtime.convT64(SB)
  main.go:23		0x10a6eae		488b442408		MOVQ 0x8(SP), AX
  main.go:23		0x10a6eb3		0f57c0			XORPS X0, X0
  main.go:23		0x10a6eb6		0f11442440		MOVUPS X0, 0x40(SP)
  main.go:23		0x10a6ebb		488d0dfea80000		LEAQ runtime.rodata+43168(SB), CX
  main.go:23		0x10a6ec2		48894c2440		MOVQ CX, 0x40(SP)
  main.go:23		0x10a6ec7		4889442448		MOVQ AX, 0x48(SP)
  print.go:274		0x10a6ecc		488b0565590c00		MOVQ os.Stdout(SB), AX
  print.go:274		0x10a6ed3		488d0de6540400		LEAQ go.itab.*os.File,io.Writer(SB), CX
  print.go:274		0x10a6eda		48890c24		MOVQ CX, 0(SP)
  print.go:274		0x10a6ede		4889442408		MOVQ AX, 0x8(SP)
  print.go:274		0x10a6ee3		488d442440		LEAQ 0x40(SP), AX
  print.go:274		0x10a6ee8		4889442410		MOVQ AX, 0x10(SP)
  print.go:274		0x10a6eed		48c744241801000000	MOVQ $0x1, 0x18(SP)
  print.go:274		0x10a6ef6		48c744242001000000	MOVQ $0x1, 0x20(SP)
  print.go:274		0x10a6eff		90			NOPL
  print.go:274		0x10a6f00		e8db9affff		CALL fmt.Fprintln(SB)
  main.go:23		0x10a6f05		488b6c2450		MOVQ 0x50(SP), BP
  main.go:23		0x10a6f0a		4883c458		ADDQ $0x58, SP
  main.go:23		0x10a6f0e		c3			RET
  main.go:22		0x10a6f0f		e86c9ffbff		CALL runtime.morestack_noctxt(SB)
  main.go:22		0x10a6f14		e967ffffff		JMP main.main(SB)
*/
