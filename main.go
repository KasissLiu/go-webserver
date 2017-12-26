package main

import (
	"ksweb/servers/dynamicServer"
	"ksweb/servers/fileServer"
	"net/http"
	"os"
	"time"
)

type Server struct {
	FileServer    *fileServer.FileServer
	DynamicServer *dynamicServer.DynamicServer
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

func init() {

	fileServer := fileServer.New("/nginx/openSourceCode/kasiss-web/build", "index.html", nil, nil)
	dynamicServer := dynamicServer.New()
	startTime := time.Now()
	httpServer = Server{fileServer, dynamicServer, startTime, 0, 0}
}

func loadConfig() {

}

func main() {

	http.ListenAndServe(":8000", &httpServer)

}
