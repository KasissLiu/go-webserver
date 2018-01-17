package dynamicServer

import (
	"net/http"
	"strings"

	"github.com/KasissLiu/go-webserver/routes"
)

type DynamicServer struct {
	Web      map[string]func(http.ResponseWriter, *http.Request)
	Api      map[string]func(http.ResponseWriter, *http.Request)
	execFunc func(http.ResponseWriter, *http.Request)
}

func (this *DynamicServer) CheckDynamic(path string) bool {
	path = strings.Trim(path, "/")

	if function, ok := web[path]; ok {
		this.execFunc = function
		return true
	}
	if function, ok := api[path]; ok {
		this.execFunc = function
		return true
	}
	return false

}
func (this *DynamicServer) DoDynamic(w http.ResponseWriter, r *http.Request) {
	this.execFunc(w, r)
}

func (this *DynamicServer) getAlias(req string) string {
	if val, ok := alias[req]; ok {
		return val
	}
	return req
}

var web map[string]func(http.ResponseWriter, *http.Request)
var api map[string]func(http.ResponseWriter, *http.Request)
var alias map[string]string

func init() {
	web = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	api = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	alias = make(map[string]string, 10)

	web = routes.Web
	api = routes.Api

}

func New() *DynamicServer {
	return &DynamicServer{web, api, nil}
}
