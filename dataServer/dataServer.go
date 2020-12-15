package main

import (
	"ObjectStorage/dataServer/heartbeat"
	"ObjectStorage/dataServer/locate"
	"ObjectStorage/dataServer/object"
	"ObjectStorage/dataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	//发送心跳
	go heartbeat.StartHeartBeat()
	//接收定位消息
	go locate.StartLocate()
	//若定位成功，则开始真正的接收和发送文件
	http.HandleFunc("/objects/", object.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		log.Fatal("listen error", err)
	}
}
