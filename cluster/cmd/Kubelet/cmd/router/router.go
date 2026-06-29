package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

func LoadRouter() *chi.Mux {
	Router := chi.NewMux()

	Router.Route("/v1/kubectlApi", func(api chi.Router) {
		api.Route("/get", func(r chi.Router) {
			fmt.Println("get works")
		})
	})

	return Router
}
