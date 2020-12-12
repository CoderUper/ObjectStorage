package es

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

func getMetadata(name string, versionId int) (meta Metadata, err error) {
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d/source", os.Getenv("ES_SERVER"), name, versionId)
	r, err := http.Get(url)
	if err != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to get %s_%d", name, versionId)
		return
	}
	//获取到文件元数据
	result, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(result, &meta)
	return
}

type hit struct {
	Source Metadata `json:"_source"`
}
type searchResult struct {
	Hits struct {
		Total int
		Hits  []hit
	}
}

func SearchLatestVersion(name string) (meta Metadata, err error) {
	url := fmt.Sprintf("http://%s/metadata/_search?q=name:%s&size=1&sort=version:desc", os.Getenv("ES_SERVER"), name)
	r, err := http.Get(url)
	if err != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("fialed to find %s latest version", name)
		return
	}
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	if len(sr.Hits.Hits) != 0 {
		meta = sr.Hits.Hits[0].Source

	}
	return
}
func GetMetadata(name string, version int) (meta Metadata, err error) {
	if version == 0 {
		return SearchLatestVersion(name)
	}
	return getMetadata(name, version)
}

func PutMetadata(name string, version int, size int64, hash string) error {
	doc := fmt.Sprintf(`{"name":%s,"version":%d,"size":%d,"hash":%s}`, name, version, size, hash)
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d?op_type=create", os.Getenv("ES_SERVER"), name, version)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(doc))
	r, err := client.Do(request)
	if err != nil {
		return err
	}
	//如果多个客户端同时上传一个元数据，只有一个成功，其他的会增加版本号继续递归上传，直到成功
	if r.StatusCode == http.StatusConflict {
		return PutMetadata(name, version+1, size, hash)
	}
	if r.StatusCode != http.StatusCreated {
		result, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("failed to put metadata : %d  %s", r.StatusCode, string(result))
	}
	return nil
}

func AddVersion(name, hash string, size int64) error {
	version, e := SearchLatestVersion(name)
	if e != nil {
		return e
	}
	return PutMetadata(name, version.Version+1, size, hash)
}

func SearchAllVersions(name string, from, size int) ([]Metadata, error) {
	//构造请求url
	url := fmt.Sprintf("http://%s//metadata/_search?sort=name,version&from=%d&size=%d", os.Getenv("ES_SERVER"), from, size)
	if name != "" {
		url += "&q=name:" + name
	}
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	metas := make([]Metadata, 0)
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	for i := range sr.Hits.Hits {
		metas = append(metas, sr.Hits.Hits[i].Source)
	}
	return metas, nil
}
