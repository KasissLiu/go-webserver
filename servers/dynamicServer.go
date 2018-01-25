package servers

import (
	"net/http"
	"strings"
)

type DynamicServer struct {
	Web      map[string]func(http.ResponseWriter, *http.Request)
	Api      map[string]func(http.ResponseWriter, *http.Request)
	execFunc func(http.ResponseWriter, *http.Request)
}

func (this *DynamicServer) CheckDynamic(path string) bool {
	path = strings.Trim(path, "/")

	if function, ok := Web[path]; ok {
		this.execFunc = function
		return true
	}
	if function, ok := Api[path]; ok {
		this.execFunc = function
		return true
	}
	return false

}
func (this *DynamicServer) DoDynamic(w http.ResponseWriter, r *http.Request) {
	this.execFunc(w, r)
}

func (this *DynamicServer) getAlias(req string) string {
	if val, ok := Alias[req]; ok {
		return val
	}
	return req
}

func NewDynamicServer() *DynamicServer {
	return &DynamicServer{Web, Api, nil}
}
