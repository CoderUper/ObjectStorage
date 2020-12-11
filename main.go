package main

import (
	"ObjectStorage/object"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/object/",object.Handler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("listen error",err)
	}
}
