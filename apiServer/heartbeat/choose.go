package heartbeat

import "math/rand"

func ChooseRandomDataServers(n int, exclude map[int]string) (ds []string) {
	candidates := make([]string, 0)
	//exclude the server which has the data
	reverseExcludeServers := make(map[string]int)
	for id, addr := range exclude {
		reverseExcludeServers[addr] = id
	}
	//get current data servers that online
	servers := GetDataServers()
	for _, s := range servers {
		//judge the server exclude or include
		_, excluded := reverseExcludeServers[s]
		if !excluded {
			candidates = append(candidates, s)
		}
	}
	length := len(candidates)
	if length < n {
		return
	}
	p := rand.Perm(length)
	for i := 0; i < n; i++ {
		ds = append(ds, candidates[p[i]])
	}
	return
}
