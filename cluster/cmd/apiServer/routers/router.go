package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tekluabayneh/gok8s/cmd/apiServer/handlers"
	internals "github.com/tekluabayneh/gok8s/internals/apiserver"
)

func LoadRouter() *chi.Mux {
	router := chi.NewMux()

	router.Route("/api/v1", func(api chi.Router) {
		/////
		etcdSore := &internals.EtcdStore{}
		Handler := &handlers.PodHnalder{Store: etcdSore} // in here i am injectin depedency of the etcd cuz hander does not know about it but when we pass it lie can get form p.store.Create or Get

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "api Server is running",
			})
		})

		api.Get("/namespace/{namespace}/pods/{name}", Handler.Get)
		api.Post("/{namespace}/{name}", Handler.Create)
	})

	return router
}
