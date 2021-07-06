package test

import (
	"go-code/study/dispatcher/heartbeat"
	"testing"
)

func TestHeartBeatSend(t *testing.T) {
	heartbeat.Start()
}

func TestHeartBeatListen(t *testing.T) {
	heartbeat.Listen()
}