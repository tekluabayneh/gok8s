package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tekluabayneh/gok8s/internals/decoder"
)

func LoadRouter() *chi.Mux {
	Router := chi.NewMux()

	// TODO
	// make sure to inject/add middleware to api.Route so they will be restricted
	Router.Route("/api/v1/kubectl", func(api chi.Router) {
		//  health chec
		api.Route("/", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("get pods")
				decoder.Encoder(w, http.StatusOK, map[string]string{
					"msg": "kubelet api server is healthy",
				})
			})
		})

		// pods — GET: List of all pods on this node (what kubelet sees)
		Router.Get("/pods", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods")
		})

		// logs/{pod}/{container} — GET: Stream container logs (what kubectl logs uses)
		// kubectl will reach out here to get logs for container
		Router.Get("/logs/{pod}/{container}", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods stream")
		})

		// exec/{pod}/{container} — POST: Execute command in container
		// with stdin/stdout/stderr (what kubectl exec uses, bidirectional over WebSocket/SPDY)
		// this is the route that help us run command inside container
		Router.Post("/exec/{pod}/{container}", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods stream")
		})

		// metrics — GET: Prometheus-format node + pod metrics
		Router.Get("/matrics", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods stream")
		})

		// stats — GET: Node and container statistics (used by metrics-server)
		Router.Get("/stats", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods stream")
		})

		// spec — GET: Node hardware spec
		Router.Get("/spec", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods stream")
		})

		// proxy/{path} — Proxy requests to other kubelet endpoints
		Router.Post("/proxy{path}", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get pods stream")
		})
	})

	return Router
}
