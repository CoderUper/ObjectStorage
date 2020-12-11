package main

import (
	"ObjectStorage/apiServer/heartbeat"
	"ObjectStorage/apiServer/locate"
	"ObjectStorage/apiServer/object"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartBeat()
	//相当于PUT上传数据
	http.HandleFunc("/objects", object.Handler)
	//相当于GET请求数据
	http.HandleFunc("/locate", locate.Handler)
	log.Fatalln(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
