package loadbalance

import (
	"errors"
	"math/rand"
	"time"
)

type Random struct {
	nodes []string
}

func (b *Random) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params is invalid")
	}
	for _, param := range params {
		b.nodes = append(b.nodes, param)
	}
	return nil
}

func (b *Random) Get() (string, error) {
	if len(b.nodes) == 0 {
		return "", errors.New("nodes is empty")
	}
	rand.Seed(time.Now().Unix())
	return b.nodes[rand.Intn(len(b.nodes))], nil
}
