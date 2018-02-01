package routes

import (
	"net/http"
)

var Web map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Web = make(map[string]func(http.ResponseWriter, *http.Request), 10)
}
