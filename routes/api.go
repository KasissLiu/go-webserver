package routes

import (
	"net/http"
)

var Api map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Api = make(map[string]func(http.ResponseWriter, *http.Request), 10)

	//Api["api/hello"] = test.TestApi
}
