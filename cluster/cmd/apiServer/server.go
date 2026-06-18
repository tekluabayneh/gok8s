package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	Client "github.com/tekluabayneh/gok8s/cmd/etcd"
)

// TODO
//	create server without any

func ApiServer() {
	etcdDB, err := Client.InitEtcd()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := etcdDB.Put(ctx, "example_key", "example_value")
	if err != nil {
		return
	}

	fmt.Println("res", res)

	fmt.Println(etcdDB.Get(ctx, "example_key"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})
	mux.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	fmt.Println("apiServer started")
	http.ListenAndServe("localhost:5000", mux)
}
