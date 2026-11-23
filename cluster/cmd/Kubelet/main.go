package main

import (
	"fmt"
	"time"

	"github.com/tekluabayneh/gok8s/cmd/Kubelet/internal"
)

var (
	defaultbool bool
	HostName    = "kind-control-planemm"
)

func SendhearBeat() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticker start hearbeat")
		}
	}
}

func main() {
	// cli, req := config.CreatApi()
	// go SendhearBeat() /// having function with ticker time and for loop and go routing are hard to maintine since they run in diffrent proccess tey each access and update the time resulting error so the ticker start sending wont' really wait the second we specify so this need to be updated
	// fmt.Println("che")
	// server.Start().KubeServer() // this server can be fork-join for our demon job

	internal.RegNode()
}
