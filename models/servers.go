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

//服务器运行数据实例
var State *ServerState = &ServerState{}

//更新运行数据
func (s *ServerState) SetServerState(m map[string]int) {
	s.AccessTimes, _ = m["accessTotal"]
	s.ActTotal, _ = m["actTotal"]
	s.StartTime = time.Unix(int64(m["startTime"]), 0).Format("2006-01-02 15:04:05")
	s.ServerTime = time.Now().Format("2006-01-02 15:04:05")
}

//获取运行数据
func (s *ServerState) GetServerState() *ServerState {
	return s
}
