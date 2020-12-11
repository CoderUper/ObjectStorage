package object

import (
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	//文件体和文件名
	status, err := storeObject(r.Body, object)
	if err != nil {
		log.Println(err)
	}
	//上传文件成功
	//返回状态码
	w.WriteHeader(status)
}
