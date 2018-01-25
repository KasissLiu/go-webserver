package models

import (
	"time"
)

type ServerState struct {
	AccessTimes int    `json:"access_times"`
	ActTotal    int    `json:"act_total"`
	StartTime   string `json:"start_time"`
	ServerTime  string `json:"server_time"`
}

var State *ServerState = &ServerState{}

func (s *ServerState) SetServerState(m map[string]int) {
	s.AccessTimes, _ = m["accessTotal"]
	s.ActTotal, _ = m["actTotal"]
	s.StartTime = time.Unix(int64(m["startTime"]), 0).Format("2006-01-02 15:04:05")
	s.ServerTime = time.Now().Format("2006-01-02 15:04:05")
}

func (s *ServerState) GetServerState() *ServerState {
	return s
}
