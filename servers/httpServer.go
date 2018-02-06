package servers

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

//server结构体 保存了一个webserver的运行设置
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
	Gzip          bool
}

// 一个实现了http.ResponseWriter接口的结构体 用来进行gzip写入
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

//实现gzipResponseWriter的写方法
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

//服务器实现函数
func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	this.AccessTimes++
	syncServerState()

	requestType := this.ParseServer(r)

	//gzip 不能用于websocket链接申请
	if requestType != "websocket" {
		//如果服务器开启gzip 并且浏览器支持 则进行gzip压缩
		if this.Gzip && strings.Contains(strings.ToLower(r.Header.Get("Accept-Encoding")), "gzip") {
			//设置响应头 告知浏览器本次请求为gzip压缩
			w.Header().Set("Content-Encoding", "gzip")
			//非常重要 如果不设置此头 浏览器将不解析gzip
			w.Header().Set("Transfer-Encoding", "chunked")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			w = gzipResponseWriter{Writer: gz, ResponseWriter: w}
		}
	}

	//解析请求
	switch requestType {
	case "file":
		this.FileServer.FileOutput(w, r)
	case "dynamic":
		this.DynamicServer.Execute(w, r)
	case "websocket":
		this.WebsocketCons++
		syncServerState()
		this.DynamicServer.Execute(w, r)
		this.WebsocketCons--
	default:
		this.FileServer.IndexOutput(w)
	}

	syncServerState()

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
	//判断是否是websocket链接
	if this.DynamicServer.CheckWebsocket(r.URL.Path) {
		return "websocket"
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
