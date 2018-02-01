package routes

import (
	"net/http"

	"github.com/KasissLiu/go-webserver/controllers/ws"
)

var Ws map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Ws = make(map[string]func(http.ResponseWriter, *http.Request), 10)

	Ws["ws/server-state"] = ws.SyncServerState
}
