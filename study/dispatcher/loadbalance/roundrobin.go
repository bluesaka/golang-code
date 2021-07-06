package loadbalance

import (
	"errors"
)

type RoundRobin struct {
	curIdx int
	nodes  []string
}

func (b *RoundRobin) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params is invalid")
	}
	for _, param := range params {
		b.nodes = append(b.nodes, param)
	}
	return nil
}

func (b *RoundRobin) Get() (string, error) {
	if len(b.nodes) == 0 {
		return "", errors.New("nodes is empty")
	}
	if b.curIdx >= len(b.nodes) {
		b.curIdx = 0
	}
	node := b.nodes[b.curIdx]
	b.curIdx++
	return node, nil
}
