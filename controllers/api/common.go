package api

import (
	"encoding/json"
	"io"
	"net/http"

	m "github.com/KasissLiu/go-webserver/models"
)

func GetServerState(w http.ResponseWriter, h *http.Request) {
	status := m.State.GetServerState()
	responseJson(0, "success", status, w)
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func responseJson(code int, msg string, data interface{}, w http.ResponseWriter) {
	resp := response{code, msg, data}
	bytes, _ := json.Marshal(resp)
	io.WriteString(w, string(bytes))
}
