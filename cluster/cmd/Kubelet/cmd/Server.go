package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tekluabayneh/gok8s/cmd/Kubelet/cmd/router"
)

type app struct {
	router http.Handler
}

func Start() *app {
	router := &app{
		router: router.LoadRouter(),
	}
	return router
}

func (han *app) KubeServer() {
	PORT := "3030"

	Serv := &http.Server{
		Addr:         ":" + PORT,
		Handler:      han.router,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("server is running")
	if err := Serv.ListenAndServe(); err != nil {
		panic(err)
	}
}
