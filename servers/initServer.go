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
var HttpServerTLS Server
var HttpConfig *loadConfig.Config

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
	//http config
	httpFileRoot, error := HttpConfig.Get("http").Get("root").String()
	if error != nil {
		panic("set default root path")
	}
	httpPort, error := HttpConfig.Get("http").Get("port").Int()
	if error != nil {
		httpPort = 8000
	}
	httpIndex, error := HttpConfig.Get("http").Get("index").String()
	if error != nil {
		httpIndex = "index.html"
	}
	var httpOpen bool
	_, error = HttpConfig.Get("http").Get("type").String()
	if error == nil {
		httpOpen = true
	} else {
		httpOpen = false
	}

	//https config
	httpsFileRoot, error := HttpConfig.Get("https").Get("root").String()
	if error != nil {
		httpsFileRoot = ""
	}
	httpsIndex, error := HttpConfig.Get("https").Get("index").String()
	if error != nil {
		httpsIndex = "index.html"
	}
	httpsPort, error := HttpConfig.Get("https").Get("port").Int()
	if error != nil {
		httpsPort = 443
	}
	httpsCrt, error := HttpConfig.Get("https").Get("crt").String()
	if error != nil {
		httpsCrt = ""
	}
	httpsKey, error := HttpConfig.Get("https").Get("key").String()
	if error != nil {
		httpsKey = ""
	}

	var httpsOpen bool
	_, error = HttpConfig.Get("https").Get("type").String()
	if error == nil {
		httpsOpen = true
	} else {
		httpsOpen = false
	}

	dynamicServer := NewDynamicServer()
	StartTime := time.Now()
	//make a http Server
	httpFileServer := NewFileServer(httpFileRoot, httpIndex, nil, nil)
	HttpServer = Server{
		FileServer:    httpFileServer,
		DynamicServer: dynamicServer,
		Port:          httpPort,
		Open:          httpOpen,
		StartTime:     StartTime,
		AccessTimes:   0,
		WebsocketCons: 0}

	//make a https Server
	httpsFileServer := NewFileServer(httpsFileRoot, httpsIndex, nil, nil)
	HttpServerTLS = Server{
		FileServer:    httpsFileServer,
		DynamicServer: dynamicServer,
		Port:          httpsPort,
		Open:          httpsOpen,
		HttpsCrt:      httpsCrt,
		HttpsKey:      httpsKey,
		StartTime:     StartTime,
		AccessTimes:   0,
		WebsocketCons: 0}
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
