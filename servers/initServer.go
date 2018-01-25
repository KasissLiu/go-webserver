package servers

import (
	"net/http"
	"time"

	"github.com/KasissLiu/go-tools/loadConfig"
	"github.com/KasissLiu/go-webserver/models"
	"github.com/KasissLiu/go-webserver/routes"
)

//初始化变量
var HttpServer Server
var HttpConfig *loadConfig.Config
var DbConfig *loadConfig.Config

var Web map[string]func(http.ResponseWriter, *http.Request)
var Api map[string]func(http.ResponseWriter, *http.Request)
var Alias map[string]string

func init() {
	initConfig()
	initServer()
	initRoute()
}

func initConfig() {
	HttpConfig = loadConfig.New("kasiss", "./config/server.ini")
}

func initServer() {
	fileRoot, error := HttpConfig.Get("root").String()
	if error != nil {
		panic("set default root path")
	}
	port, error := HttpConfig.Get("port").Int()
	if error != nil {
		port = 8000
	}
	index, error := HttpConfig.Get("index").String()
	if error != nil {
		index = "index.html"
	}
	fileServer := NewFileServer(fileRoot, index, nil, nil)
	dynamicServer := NewDynamicServer()
	StartTime := time.Now()
	HttpServer = Server{fileServer, dynamicServer, port, StartTime, 0, 0}
}

func initRoute() {
	Web = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	Api = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	Alias = make(map[string]string, 10)

	Web = routes.Web
	Api = routes.Api

}

func syncServerState() {
	state := HttpServer.GetAccessStat()
	models.State.SetServerState(state)
}
