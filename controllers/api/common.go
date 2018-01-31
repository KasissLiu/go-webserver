package api

import (
	"encoding/json"
	"io"
	"net/http"

	m "github.com/KasissLiu/go-webserver/models"
)

//获取服务器的运行统计
func GetServerState(w http.ResponseWriter, r *http.Request) {
	status := m.State.GetServerState()
	responseJson(0, "success", status, w)
}

//响应数据结构体
type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//统一格式化json方法
func responseJson(code int, msg string, data interface{}, w http.ResponseWriter) {
	resp := response{code, msg, data}
	bytes, _ := json.Marshal(resp)
	io.WriteString(w, string(bytes))
}
