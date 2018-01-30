package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KasissLiu/go-webserver/servers"
)

func main() {
	processEnd := make(chan bool)
	//开启多个线程 做不同的事情
	go startWebServer()
	go startWebServerTLS()

	<-processEnd
	println("server done...")

}

//启动一个http服务器
func startWebServer() {

	if servers.HttpServer.Open {
		httpPort := ":" + strconv.Itoa(servers.HttpServer.Port)

		log.Println("web server start on port" + httpPort)
		err := http.ListenAndServe(httpPort, &servers.HttpServer)
		if err != nil {
			log.Fatal(err)
		}
	}

}

//启动一个https服务器
func startWebServerTLS() {
	if servers.HttpServerTLS.Open {
		httpsPort := ":" + strconv.Itoa(servers.HttpServerTLS.Port)
		crt := servers.HttpServerTLS.HttpsCrt
		key := servers.HttpServerTLS.HttpsKey
		log.Println("web server start on port" + httpsPort)
		err := http.ListenAndServeTLS(httpsPort, crt, key, &servers.HttpServerTLS)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//启动一个websocket服务器
func startWebSocketServer(c chan bool) {
	//do someting
}
