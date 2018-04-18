package routes

import (
	"net/http"

	"github.com/kasiss-liu/go-webserver/controllers/api"
)

var Api map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Api = make(map[string]func(http.ResponseWriter, *http.Request), 10)

	Api["api/state"] = api.GetServerState
	Api["api/user"] = api.User.GetUserById
	Api["api/userall"] = api.User.GetUserAll
}
