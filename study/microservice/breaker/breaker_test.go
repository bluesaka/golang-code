package breaker

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// TestBreaker
// exec `go test -v -run ^TestBreaker$` command to test
func TestBreaker(t *testing.T) {
	strategyOpt := StrategyOption{
		Strategy:      StrategyFail,
		FailThreshold: 2,
	}
	breaker := NewBreaker(WithName("my-breaker"), WithStrategyOption(strategyOpt))
	for i := 0; i < 20; i++ {
		log.Println("i:", i)
		breaker.Call(func() error {
			if i <= 2 || i >= 8 {
				return nil
			} else {
				return errors.New("error")
			}
		})
		fmt.Println()
		time.Sleep(time.Second)
	}
}

// exec `go test -v -run ^TestBreaker2$` command to test
func TestBreaker2(t *testing.T) {
	strategyOpt := StrategyOption{
		Strategy:      StrategyFail,
		FailThreshold: 2,
	}
	breaker := NewBreaker(WithName("my-breaker"), WithStrategyOption(strategyOpt))
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				log.Println("j:", j)
				breaker.Call(func() error {
					if j <= 2 || j >= 8 {
						return nil
					} else {
						return errors.New("error")
					}
				})
				fmt.Println()
				time.Sleep(time.Second)
			}
		}()
	}
	wg.Wait()
}
