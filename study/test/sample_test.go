/**
单元测试
goland中，右键 -> Generate -> Tests for ... 生成xxx_test.go文件

文件和函数命名规则：
- 文件必须是*_test.go
- 功能测试函数名必须以Test开头，函数参数必须是*testing.T
- 性能测试函数名必须以Benchmark开头，函数参数必须是*testing.B
- 示例测试函数名必须以Example开头，函数参数无要求

// 功能函数测试命令
go test -v sample_test.go sample.go

// 会运行Test_echo开头的函数，如Test_echo Test_echo2等
go test -v -run Test_echo sample_test.go sample.go

// 功能测试函数名必须以Test开头，testing: warning: no tests to run
go test -v -run My_Test_echo sample_test.go sample.go

// 性能测试
go test -bench=. -benchmem -run=BenchmarkEcho
go test -bench=BenchmarkEcho -benchmem -run=BenchmarkEcho

 - bench regexp 执行相应的 benchmarks，例如 -bench=.；
 - cover 开启测试覆盖率；
 - run regexp 只运行 regexp 匹配的函数，例如 -run=Test_A 那么就执行包含有 Test_A 开头的函数；
 - v 显示测试的详细命令。

需要加上sample.go，因为调用了sample.go里的相关方法，不加上会报undefined错误

*/
package test

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_echo(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"echo_test_1",
			args{
				"echo_a",
			},
			"echo_a",
		},
		{
			"echo_test_2",
			args{
				123,
			},
			123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := echo(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("echo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_echo2(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"echo_test_1",
			args{
				"echo_a",
			},
			"echo_a",
		},
		{
			"echo_test_2",
			args{
				123,
			},
			123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := echo(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("echo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func My_Test_echo(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"echo_test_1",
			args{
				"echo_a",
			},
			"echo_a",
		},
		{
			"echo_test_2",
			args{
				123,
			},
			123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := echo(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("echo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_say(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"say_test_1",
			args{
				"say_a",
			},
			"say_a",
		},
		{
			"say_test_2",
			args{
				"say_b",
			},
			"say_b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := say(tt.args.s); got != tt.want {
				t.Errorf("say() = %v, want %v", got, tt.want)
			}
		})
	}
}

/**
$ go test -v -bench=. -benchmem -run=BenchmarkEcho
$ go test -v -bench=BenchmarkEcho -benchmem -run=BenchmarkEcho
goos: darwin
goarch: amd64
pkg: go-code/study/test
BenchmarkEcho
BenchmarkEcho-8         1000000000               0.272 ns/op           0 B/op          0 allocs/op
PASS
ok      go-code/study/test      3.276s
 */
func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo(i)
	}
}

/**
$ go test -v -bench=BenchmarkSay -benchmem -run=BenchmarkSay
goos: darwin
goarch: amd64
pkg: go-code/study/test
BenchmarkSay
BenchmarkSay-8             16072            126426 ns/op           71267 B/op      15972 allocs/op
PASS
ok      go-code/study/test      3.100s
 */
func BenchmarkSay(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < b.N; i++ {
				say(strconv.Itoa(i))
			}
		}
	})
}