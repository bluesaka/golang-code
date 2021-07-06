package loadbalance

import (
	"errors"
	"strconv"
)

type WeightRoundRobin struct {
	curIdx     int
	nodes      []WeightNode
	nodeSeries []string
}

type WeightNode struct {
	node   string
	weight int
}

func (b *WeightRoundRobin) Add(params ...string) error {
	if len(params) < 2 || len(params)&1 == 1 {
		return errors.New("params length must be even")
	}

	for i := 0; i < len(params); i += 2 {
		weight, err := strconv.Atoi(params[i+1])
		if err != nil {
			return err
		}
		b.nodes = append(b.nodes, WeightNode{
			node:   params[i],
			weight: weight,
		})
	}

	b.buildNodeSeries()
	return nil
}

func (b *WeightRoundRobin) Get() (string, error) {
	if len(b.nodes) == 0 {
		return "", errors.New("nodes is empty")
	}
	if b.curIdx >= len(b.nodeSeries) {
		b.curIdx = 0
	}
	node := b.nodeSeries[b.curIdx]
	b.curIdx++
	return node, nil
}

func (b *WeightRoundRobin) buildNodeSeries() {
	b.nodeSeries = nil
	nodes := make([]WeightNode, len(b.nodes))
	copy(nodes, b.nodes)

	for {
		weight, sum, index := 0, 0, 0
		for i, node := range nodes {
			sum += node.weight
			if node.weight > weight {
				weight = node.weight
				index = i
			}
		}
		if sum == 0 {
			break
		}
		nodes[index].weight--
		b.nodeSeries = append(b.nodeSeries, nodes[index].node)
	}

}
