package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KasissLiu/go-webserver/servers"
)

func main() {
	httpPort := ":" + strconv.Itoa(servers.HttpServer.Port)
	log.Println("web server start on port" + httpPort)
	err := http.ListenAndServe(httpPort, &servers.HttpServer)
	if err != nil {
		log.Fatal("web server start failed")
	}

}
