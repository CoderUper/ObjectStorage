package temp

import (
	"ObjectStorage/dataServer/locate"
	"ObjectStorage/src/lib/utils"
	"compress/gzip"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func (t *tempInfo) hash() string {
	return strings.Split(t.Name, ".")[0]
}

func (t *tempInfo) id() int {
	s := strings.Split(t.Name, ".")
	id, _ := strconv.Atoi(s[1])
	return id
}

func commitTempObject(datFile string, tempinfo *tempInfo) {
	f, _ := os.Open(datFile)
	shardHash := url.PathEscape(utils.CalculateHash(f))
	defer f.Close()
	f.Seek(0, io.SeekStart)
	w, _ := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + tempinfo.Name + "." + shardHash)
	w2 := gzip.NewWriter(w)
	io.Copy(w2, f)
	w2.Close()
	locate.Add(tempinfo.hash(), tempinfo.id())
	os.Remove(datFile)
	os.Remove(strings.Split(datFile, ".")[0])
}
