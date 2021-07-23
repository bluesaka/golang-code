package main

import (
	"go-code/study/microservice/configcenter/nacos"
	"log"
)

func main() {
	go nacos.InitNacos()

	log.Println("start...")
	select {}
}
