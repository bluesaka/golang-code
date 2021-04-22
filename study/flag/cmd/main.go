package main

import (
	"flag"
	"log"
)

/**
命令行参数解析，以下方式都支持
go run main.go -f1
go run main.go -f1 true
go run main.go -f1=true
go run main.go --f1 true
go run main.go --f1=true

*/
var f1 = flag.Bool("f1", false, "fl value")
var f2 = flag.String("f2", "default f2", "f2 value")

func main() {
	flag.Parse()
	log.Printf("f1: %v, f2: %v\n", *f1, *f2)
}
