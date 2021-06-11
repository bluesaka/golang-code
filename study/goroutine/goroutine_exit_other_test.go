package goroutine

import (
	"runtime"
	"testing"
	"time"
)

/**
$ go test -v -run ^TestDo$
=== RUN   TestDo
    goroutine_exit_other_test.go:10: goroutine num: 2
    goroutine_exit_other_test.go:13: goroutine num: 3
--- PASS: TestDo (2.27s)
PASS
ok      go-code/study/goroutine 2.286s
 */
func TestDo(t *testing.T) {
	t.Log("goroutine num:", runtime.NumGoroutine())
	SendTask()
	time.Sleep(time.Second)
	t.Log("goroutine num:", runtime.NumGoroutine())
}

/**
$ go test -v -run ^TestDoWithClose$
=== RUN   TestDoWithClose
    goroutine_exit_other_test.go:29: goroutine num: 2
2021/06/09 15:13:42 chan closed
    goroutine_exit_other_test.go:33: goroutine num: 2
--- PASS: TestDoWithClose (2.25s)
PASS
ok      go-code/study/goroutine 2.265s
 */
func TestDoWithClose(t *testing.T) {
	t.Log("goroutine num:", runtime.NumGoroutine())
	SendTaskWithClose()
	time.Sleep(time.Second)
	t.Log("goroutine num:", runtime.NumGoroutine())
}
