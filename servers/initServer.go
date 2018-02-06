package servers

import (
	"net/http"
	"time"

	"github.com/KasissLiu/go-tools/loadConfig"
	"github.com/KasissLiu/go-webserver/models"
	"github.com/KasissLiu/go-webserver/routes"
)

//初始化变量
var HttpServer Server             //http实例
var HttpServerTLS Server          //https实例
var HttpConfig *loadConfig.Config //服务器配置实例

//路由配置
var Web map[string]func(http.ResponseWriter, *http.Request)
var Api map[string]func(http.ResponseWriter, *http.Request)
var Ws map[string]func(http.ResponseWriter, *http.Request)
var Alias map[string]string

//初始化函数
func init() {
	initConfig()
	initRoute()
	initServer()
}

//初始化服务器配置
func initConfig() {
	HttpConfig = loadConfig.New("kasiss", "./config/server.ini")
}

//初始化路由数据
func initRoute() {
	Web = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	Api = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	Ws = make(map[string]func(http.ResponseWriter, *http.Request), 0)
	Alias = make(map[string]string, 10)

	Web = routes.Web
	Api = routes.Api
	Ws = routes.Ws
}

//初始化服务器实例
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
	//服务器配置gzip压缩
	var gzip bool
	gzipOpen, error := HttpConfig.Get("server").Get("gzip").String()
	if error != nil {
		gzipOpen = "off"
	}
	if gzipOpen == "on" {
		gzip = true
	} else {
		gzip = false
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
		WebsocketCons: 0,
		Gzip:          gzip,
	}

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
		WebsocketCons: 0,
		Gzip:          gzip,
	}
}

//主动同步服务器数据
func syncServerState() {
	state := HttpServer.GetAccessStat()
	models.State.SetServerState(state)
}
