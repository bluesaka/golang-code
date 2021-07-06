package loadbalance

const (
	TypeRandom = iota + 1
	TypeRoundRobin
	TypeWeightRoundRobin
)

func Factory(t int) LoadBalance {
	switch t {
	case TypeRandom:
		return new(Random)
	case TypeRoundRobin:
		return new(RoundRobin)
	case TypeWeightRoundRobin:
		return new(WeightRoundRobin)
	default:
		return new(Random)
	}
}
