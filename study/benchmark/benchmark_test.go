package benchmark

import (
	"bytes"
	"sync"
	"testing"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	data = make([]byte, 10000)
)

/**
基准测试：
文件名为 xxx_test.go，必须以 `_test` 结尾
函数名为 BenchmarkXXX

go test -bench=. -benchmem
go test -bench=^BenchmarkBufferWithPool$ -benchmem

BenchmarkBufferWithPool-8        9974926               121 ns/op               0 B/op          0 allocs/op
BenchmarkBuffer-8                1004248              1222 ns/op           10240 B/op          1 allocs/op

 */
func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

/**

BenchmarkBufferWithPool-8        9974926               121 ns/op               0 B/op          0 allocs/op
BenchmarkBuffer-8                1004248              1222 ns/op           10240 B/op          1 allocs/op

 */
func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}


