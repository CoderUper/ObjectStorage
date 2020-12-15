package locate

import (
	"ObjectStorage/src/lib/rabbitmq"
	"ObjectStorage/src/lib/rs"
	"ObjectStorage/src/lib/types"
	"encoding/json"
	"os"
	"time"
)

func Locate(name string) map[int]string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(1 * time.Second)
		q.Close()
	}()

	locateInfo := make(map[int]string)
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return locateInfo
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return locateInfo
}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS
}
