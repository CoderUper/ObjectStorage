package object

import (
	"ObjectStorage/src/lib/es"
	"ObjectStorage/src/lib/utils"
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	//1.获取请求头哈希
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash code in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//以唯一的哈希值作为文件名进行存储
	status, err := storeObject(r.Body, hash)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}
	//3.获取真正的文件名，再元数据服务器上存储最新的元数据
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	//文件体和文件名
	size := utils.GetSizeFromHeader(r.Header)
	err = es.AddVersion(name, hash, size)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

	}
	return

	//上传文件成功
	//返回状态码
	w.WriteHeader(status)
}
