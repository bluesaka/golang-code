/**
unsafe包

type ArbitraryType int
type Pointer *ArbitraryType

func Sizeof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Alignof(x ArbitraryType) uintptr

Sizeof 返回类型x所占据的字节数，不包含x指向内容的大小
Offsetof 返回结构体成员在内存中的位置离结构体起始处的字节数，所传参数必须是结构体成员
Alignof 返回m，m是指当类型进行内存对骑士，它分配的内存地址能整除m

任何类型的指针 和 unsafe.Pointer 可以相互转换
uintptr类型 和 unsafe.Pointer 可以相互转换

uintptr <-> unsafe.Pointer <-> *T
*/

package unsafe

import (
	"fmt"
	"unsafe"
)

func SliceTest() {
	// runtime/slice.go
	// slice的结构体定义
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	// 调用make函数创建slice，底层调用makeslice函数，返回slice结构体
	// func makeslice(et *_type, len, cap int) slice
	s := make([]int, 6, 15)

	// 转换过程: &s => pointer => uintptr => pointer => *int => int
	Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 6 6

	Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 15 15
}

func MapTest() {
	// runtime/map.go
	type hmap struct {
		// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
		// Make sure this stays in sync with the compiler's definition.
		count     int // # live cells == size of map.  Must be first (used by len() builtin)
		flags     uint8
		B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
		noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
		hash0     uint32 // hash seed

		buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
		oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
		nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

		//extra *mapextra // optional fields
	}

	// 调用make函数创建map，底层调用makemap函数，返回hmap的指针
	// func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	// 转换过程：&m => pointer => **int => int
	Len := **(**int)(unsafe.Pointer(&m))
	fmt.Println(Len, len(m)) // 2 2
}
