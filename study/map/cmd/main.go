package main

func main() {
	//_map.MapGCIssue()
	//_map.MapRange()
	MapGCTest()
}

/**
go tool compile -S -N -l main.go | grep CALL

0x0020 00032 (main.go:5)        CALL    "".MapGCTest(SB)
0x002e 00046 (main.go:3)        CALL    runtime.morestack_noctxt(SB)
0x0051 00081 (main.go:9)        CALL    runtime.makemap(SB)
0x0079 00121 (main.go:10)       CALL    runtime.mapassign_fast64(SB)
0x00cf 00207 (main.go:11)       CALL    runtime.mapassign_fast64(SB)
0x0120 00288 (main.go:12)       CALL    runtime.mapassign_fast64(SB)
0x016e 00366 (main.go:14)       CALL    runtime.mapdelete_fast64(SB)
0x0184 00388 (main.go:12)       CALL    runtime.gcWriteBarrier(SB)
0x0192 00402 (main.go:11)       CALL    runtime.gcWriteBarrier(SB)
0x01a3 00419 (main.go:10)       CALL    runtime.gcWriteBarrier(SB)
0x01ad 00429 (main.go:8)        CALL    runtime.morestack_noctxt(SB)

 */
func MapGCTest() {
	m := make(map[int]string, 9)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	delete(m, 1)
}
