package goroutine

import "testing"

/**
$ go test -v -run ^TestTimeout$
=== RUN   TestTimeout
    goroutine_exit_timeout_test.go:9: goroutine num: 1002
--- PASS: TestTimeout (3.32s)
PASS
ok      go-code/study/goroutine 3.338s
 */
func TestTimeout(t *testing.T) {
	test(t, DoSleep)
}

/**
$ go test -v -run ^TestTimeoutWithBuffer$
=== RUN   TestTimeoutWithBuffer
    goroutine_exit_timeout_test.go:21: goroutine num: 2
--- PASS: TestTimeoutWithBuffer (3.28s)
PASS
ok      go-code/study/goroutine 3.296s
 */
func TestTimeoutWithBuffer(t *testing.T) {
	testWithBuffer(t, DoSleep)
}

/**
$ go test -v -run ^TestTimeoutWithSelect$
=== RUN   TestTimeoutWithSelect
    goroutine_exit_timeout_test.go:31: goroutine num: 2
--- PASS: TestTimeoutWithSelect (3.29s)
PASS
ok      go-code/study/goroutine 3.305s
 */
func TestTimeoutWithSelect(t *testing.T) {
	test(t, DoSleepWithSelect)
}

/**
$ go test -v -run ^TestTimeout2Phases$
=== RUN   TestTimeout2Phases
    goroutine_exit_timeout.go:119: goroutine num: 2
--- PASS: TestTimeout2Phases (4.31s)
PASS
ok      go-code/study/goroutine 4.320s
 */
func TestTimeout2Phases(t *testing.T) {
	test2PhasesTimeout(t)
}

