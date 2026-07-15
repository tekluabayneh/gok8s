package main

import (
	"fmt"
	"time"

	server "github.com/tekluabayneh/gok8s/cmd/Kubelet/cmd"
)

func SendhearBeat() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticker start hearbeat")
		default:
			fmt.Println("not yet")
		}
	}
}

func main() {
	go SendhearBeat() /// having function with ticker time and for loop and go routing are hard to maintine since they run in diffrent proccess tey each access and update the time resulting error so the ticker start sending wont' really wait the second we specify so this need to be updated

	server.Start().KubeServer() // this server can be fork-join for our demon job
}
