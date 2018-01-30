package servers

import (
	"net/http"
	"os"
	"time"
)

type Server struct {
	FileServer    *FileServer
	DynamicServer *DynamicServer
	Port          int
	Open          bool
	HttpsCrt      string
	HttpsKey      string
	StartTime     time.Time
	AccessTimes   int
	WebsocketCons int
}

//服务器实现函数
func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	this.AccessTimes++
	syncServerState()
	//解析请求
	switch this.ParseServer(r) {
	case "file":
		this.FileServer.FileOutput(w, r)
	case "dynamic":
		this.DynamicServer.DoDynamic(w, r)
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
func (this *Server) GetAccessStat() map[string]int {
	stat := make(map[string]int, 0)
	stat["accessTotal"] = this.AccessTimes
	stat["actTotal"] = this.WebsocketCons
	stat["startTime"] = int(this.StartTime.Unix())
	return stat
}
