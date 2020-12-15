package object

import (
	"ObjectStorage/apiServer/heartbeat"
	"ObjectStorage/apiServer/locate"
	_ "ObjectStorage/src/lib/objectstream"
	"ObjectStorage/src/lib/rs"
	"fmt"
	_ "io"
)

func GetStream(hash string, size int64) (*rs.RSGetStream, error) {
	locateInfo := locate.Locate(hash)
	//data  loss completely
	if len(locateInfo) < rs.DATA_SHARDS {
		return nil, fmt.Errorf("object %s locate fail,result is %v", hash, locateInfo)
	}
	//data loss partly,choose some other data servers
	dataServers := make([]string, 0)
	if len(locateInfo) != rs.ALL_SHARDS {
		dataServers = heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS-len(locateInfo), locateInfo)
	}
	return rs.NewRSGetStream(locateInfo, dataServers, hash, size)
}
