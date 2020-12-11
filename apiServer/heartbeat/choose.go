package heartbeat

import "math/rand"

func ChooseRandomDataServer() string {
	servers := GetDataServers()
	if len(servers) == 0 {
		return ""
	}
	return servers[rand.Intn(len(servers))]
}
