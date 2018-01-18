package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/KasissLiu/go-webserver/servers/dynamicServer"
	"github.com/KasissLiu/go-webserver/servers/fileServer"
	"github.com/KasissLiu/tools/loadConfig"
)

type Server struct {
	FileServer    *fileServer.FileServer
	DynamicServer *dynamicServer.DynamicServer
	port          int
	startTime     time.Time
	accessTimes   int
	websocketCons int
}

//服务器实现函数
func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	this.accessTimes++

	//解析请求
	switch this.ParseServer(r) {
	case "file":
		this.FileServer.FileOutput(w, r)
	case "dynamic":
		this.DynamicServer.DoDynamic(w, r)
	case "stat":

	default:
		this.FileServer.IndexOutput(w)
	}

}

//解析请求类型 判断是否是文件请求或者数据请求
func (this *Server) ParseServer(r *http.Request) string {
	//判断是否是静态文件
	if this.checkFileExist(this.FileServer.BaseDir + r.URL.Path) {
		return "file"
	}
	//判断是否是动态地址
	//return "dynamic"
	if this.DynamicServer.CheckDynamic(r.URL.Path) {
		return "dynamic"
	}

	return "index"

}

//判断文件是否存在
func (this *Server) checkFileExist(path string) bool {
	//获取文件/文件夹信息
	f, err := os.Stat(path)
	//如果存在 则判断是否是文件夹
	if err == nil && !f.IsDir() {
		return true
	}
	//不存在或者不是文件 返回false
	return false
}

//服务器统计
func (this *Server) getAccessStat() (stat map[string]int) {
	stat["accessTotal"] = this.accessTimes
	stat["actTotal"] = this.websocketCons
	stat["startTime"] = int(this.startTime.Unix())
	return stat
}

var httpServer Server
var httpConfig *loadConfig.Config

func init() {
	initConfig()
	initServer()
}

func initConfig() {
	httpConfig = loadConfig.New("kasiss", "./config/server.ini")
}

func initServer() {
	fileRoot, error := httpConfig.Get("root").String()
	if error != nil {
		panic("set default root path")
	}
	port, error := httpConfig.Get("port").Int()
	if error != nil {
		port = 8000
	}
	index, error := httpConfig.Get("index").String()
	if error != nil {
		index = "index.html"
	}
	fileServer := fileServer.New(fileRoot, index, nil, nil)
	dynamicServer := dynamicServer.New()
	startTime := time.Now()
	httpServer = Server{fileServer, dynamicServer, port, startTime, 0, 0}
}

func main() {
	httpPort := ":" + strconv.Itoa(httpServer.port)
	http.ListenAndServe(httpPort, &httpServer)

}
