package test

import (
	"fmt"
	"go-code/study/dispatcher/loadbalance"
	"testing"
)

func TestWeightRoundRobin(t *testing.T) {
	lb := loadbalance.Factory(loadbalance.TypeWeightRoundRobin)
	lb.Add("a", "1", "b", "2")
	for i := 0; i < 10; i++ {
		fmt.Println(lb.Get())
	}
	fmt.Println()

	lb.Add("c", "4")
	for i := 0; i < 10; i++ {
		fmt.Println(lb.Get())
	}
}
