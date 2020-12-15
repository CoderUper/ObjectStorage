package object

import (
	"ObjectStorage/dataServer/locate"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func get(w http.ResponseWriter, r *http.Request) {
	file := getFile(strings.Split(r.URL.EscapedPath(), "/")[2])
	if file == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sendFile(w, file)
}

func getFile(name string) string {
	fileName := os.Getenv("STORAGE_ROOT") + "/objects/" + name + ".*"
	files, _ := filepath.Glob(fileName)
	if len(files) != 1 {
		locate.Del(strings.Split(name, ".")[0])
		return ""
	}
	file := files[0]
	h := sha256.New()
	sendFile(h, file)
	shardHash := url.PathEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	hash := strings.Split(file, ".")[2]
	if shardHash != hash {
		log.Println("object hash mismatched, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}

func sendFile(w io.Writer, file string) {
	f, e := os.Open(file)
	if e != nil {
		log.Println(e)
		return
	}
	defer f.Close()
	gzipStream, e := gzip.NewReader(f)
	if e != nil {
		log.Println(e)
		return
	}
	io.Copy(w, gzipStream)
	gzipStream.Close()
}

func del(w http.ResponseWriter, r *http.Request) {
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/" + hash + ".*")
	if len(files) != 1 {
		return
	}
	locate.Del(hash)
	os.Rename(files[0], os.Getenv("STORAGE_ROOT")+"/garbage/"+filepath.Base(files[0]))
}
