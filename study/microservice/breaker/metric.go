package breaker

import "time"

type Metric struct {
	WindowBatch       uint64
	WindowStartTime   time.Time
	CountAll          uint64
	CountSuccess      uint64
	CountFail         uint64
	ContinuousSuccess uint64
	ContinuousFail    uint64
}

func (m *Metric) NewWindowBatch() {
	m.WindowBatch++
}

func (m *Metric) onSuccess() {
	m.CountAll++
	m.CountSuccess++
	m.ContinuousSuccess++
	m.CountFail = 0
}

func (m *Metric) onFail() {
	m.CountAll++
	m.CountFail++
	m.ContinuousFail++
	m.ContinuousSuccess = 0
}

func (m *Metric) OnReset() {
	m.CountAll = 0
	m.CountSuccess = 0
	m.CountFail = 0
	m.ContinuousSuccess = 0
	m.ContinuousFail = 0
}
