package object

import (
	"ObjectStorage/src/lib/es"
	"log"
	"net/http"
	"strings"
)

func del(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	//获取文件的最新元数据
	meta, err := es.SearchLatestVersion(name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//获取元数据成功，添加新的元数据，标识此数据被删除
	err = es.PutMetadata(name, meta.Version+1, 0, "")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
