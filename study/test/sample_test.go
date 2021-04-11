/**
单元测试
goland中，右键 -> Generate -> Tests for ... 生成xxx_test.go文件

测试命令：
go test -v sample_test.go sample.go
go test -v -run Test_say sample_test.go sample.go

 - bench regexp 执行相应的 benchmarks，例如 -bench=.；
 - cover 开启测试覆盖率；
 - run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
 - v 显示测试的详细命令。

需要加上sample.go，因为调用了sample.go里的相关方法，不加上会报undefined错误

*/
package test

import (
	"reflect"
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
