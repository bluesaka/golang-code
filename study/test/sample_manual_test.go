/**
单元测试
assert断言

go test -v sample.go sample_manual_test.go
go test -v -run Test_echo_manual sample.go sample_manual_test.go
*/
package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_echo_manual(t *testing.T) {
	assert.Equal(t, echo("echo test"), "echo test")
	assert.Equal(t, echo(678), 678)
}

func Test_say_manual(t *testing.T) {
	assert.Equal(t, say("say test"), "say test")
	assert.Equal(t, say("say test2"), "say test2")
}
