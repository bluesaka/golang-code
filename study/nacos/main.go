package main

import (
	"go-code/study/nacos/config"
	"log"
)

func main() {
	go config.InitNacos()

	log.Println("start...")
	select {}
}
