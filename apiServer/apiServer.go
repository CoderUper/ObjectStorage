package main

import (
	"ObjectStorage/apiServer/heartbeat"
	"ObjectStorage/apiServer/locate"
	"ObjectStorage/apiServer/object"
	"ObjectStorage/apiServer/temp"
	"ObjectStorage/apiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartBeat()
	//相当于PUT上传数据
	http.HandleFunc("/objects/", object.Handler)
	//相当于GET请求数据
	http.HandleFunc("/locate/", locate.Handler)
	//Get版本所有信息，直接查询元数据服务器
	http.HandleFunc("/versions/", versions.Handler)
	//
	http.HandleFunc("/temp/", temp.Handler)

	log.Fatalln(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
