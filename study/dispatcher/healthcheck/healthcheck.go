package healthcheck

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	addrList      = make([]string, 0)
	failCountList = make(map[string]int)
	maxRetryTimes = 3
	mu            sync.Mutex
)

func Start() {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		check()
	}
}

func check() {
	for i, addr := range addrList {
		if pingCheck(addr) {
			log.Printf(" addr: %s ping ok", addr)
			continue
		}
		mu.Lock()
		failCountList[addr]++
		if failCountList[addr] > maxRetryTimes {
			addrList = append(addrList[:i], addrList[i+1:]...)
			log.Printf("addr: %s removed\n", addr)
		}
		mu.Unlock()
	}
}

func pingCheck(addr string) bool {
	resp, err := http.Get(addr)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func AddAddr(addr ...string) {
	mu.Lock()
	defer mu.Unlock()
	addrList = append(addrList, addr...)
}

func GetAliveAddrList() []string {
	mu.Lock()
	defer mu.Unlock()
	return addrList
}
