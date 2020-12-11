package object

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const STORAGE_ROOT string = "E:\\Goland\\src\\ObjectStorage\\ObjectsRoot"
func Handler(w http.ResponseWriter,r * http.Request){
	m := r.Method
	if m==http.MethodPut{
		put(w,r)
		return
	}
	if m==http.MethodGet{
		get(w,r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func put(w http.ResponseWriter,r *http.Request){
	f,err := os.Create(STORAGE_ROOT+"\\objects\\"+strings.Split(r.URL.EscapedPath(),"/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if _,err = io.Copy(f,r.Body);err!=nil{
		log.Fatalln(err)
	}
}

func get(w http.ResponseWriter,r *http.Request){
	f,err := os.Open(STORAGE_ROOT+"\\objects\\"+strings.Split(r.URL.EscapedPath(),"/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w,f)
}
