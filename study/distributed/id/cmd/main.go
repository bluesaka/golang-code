package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"go-code/study/distributed/id"
	"time"
)

func main() {
	s1()
	//s2()
}

func s1() {
	s, _ := id.NewSnowflake(1)
	i := 0
	go func() {
		a := 5000000
		for a >= 0 {
			//fmt.Println(s.Generate())
			s.Generate()
			a--
			i++
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(i)
}

func s2() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	i := 0
	go func() {
		for {
			fmt.Println(node.Generate())
			i++
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(i)
}
