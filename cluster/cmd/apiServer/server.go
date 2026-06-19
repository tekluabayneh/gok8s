package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tekluabayneh/gok8s/cmd/apiServer/routers"
)

// Does this data live on the stack or the heap?
// Am I passing a value, a pointer, or a reference—and who owns it?
// Can multiple threads/goroutines touch this at the same time?
// What is the source of truth—and where is the state machine?
// When this loop/function panics or errors out, what state is left behind?
// What is the absolute worst-case time and space complexity here?
// What system call does this trigger under the hood?
// What happens to the file descriptor, socket, or pipe if this process dies?
// Why this specific data structure shape and not another?
// What happens if the network latency spikes or a timeout occurs?
type APIServerType struct {
	router http.Handler
}

func AppAPIServerNew() *APIServerType {
	return &APIServerType{
		router: routers.LoadRouter(),
	}
}

func (app *APIServerType) APIServerStart() {
	PORT := "5000"
	apiServ := &http.Server{
		Addr:         ":" + PORT,
		Handler:      app.router,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Printf("Server is running on post %v", PORT)
	if err := apiServ.ListenAndServe(); err != nil {
		log.Fatal("Server failed to run ", err.Error())
	}
}
