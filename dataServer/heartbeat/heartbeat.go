package heartbeat

import (
	"ObjectStorage/src/lib/rabbitmq"
	"os"
	"time"
)

func StartHeartBeat() {
	q := rabbitmq.New("RABBITMQ_SERVER")
	defer q.Close()
	for {
		q.Publish("apiServer", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
