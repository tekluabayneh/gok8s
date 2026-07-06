package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tekluabayneh/gok8s/cmd/apiServer/handlers"
	internals "github.com/tekluabayneh/gok8s/internals/apiserver"
	"github.com/tekluabayneh/gok8s/internals/etcd"
)

func LoadRouter() *chi.Mux {
	router := chi.NewMux()
	// LIFECYCLE: Main router instance. Alive for the entire lifespan of the application process.
	// MEMORY: Escapes to the heap; returned as a pointer (`*chi.Mux`) to the main server initializer.
	// FLOW: Sequential execution on the main startup thread. Used as a global read-only route map later.
	//
	// 1. Allocation: Heap. Dynamic multi-node routing struct layout.
	// 2. Concurrency: Built on a single thread; concurrently read by thousands of HTTP request goroutines safely.
	// 3. Layout: A pointer wrapper around Chi's root Radix Tree node structure.
	// 4. Failure: If initialization here fails, the process crashes immediately at startup.

	router.Route("/api/v1", func(api chi.Router) {
		// LIFECYCLE: Sub-router mount configuration scope. Runs once to construct the route branches.
		// MEMORY: Heap. Allocates a new internal sub-Mux node linked directly to the parent tree root.
		// FLOW: Executes a callback function to evaluate and append child patterns to the path prefix tree.
		//
		// 1. Allocation: Heap. Allocates memory for string path slices and endpoint handler closures.
		// 2. Concurrency: Single-threaded configuration block. Thread-safe concurrent read lookups at runtime.
		// 3. Layout: Sub-tree instance matching path prefixes.
		// 4. Failure: A panic inside this inline setup block kills the process before the server starts listening.
		//
		cli, err := etcd.InitEtcd()
		if err != nil {
			panic(fmt.Sprintf("failed to initialize etcd infrastructure: %v", err))
		}

		etcdStore := &internals.EtcdStore{
			Client: cli,
		}

		// LIFECYCLE: Shared infrastructure state. Persistent database connection context wrapper.
		// MEMORY: Heap allocation via pointer referencing (`&`). Escapes because it's wrapped into the handler.
		// FLOW: Declared once on the startup thread; acts as a global pointer pass-through to database operations.
		//
		// 1. Allocation: Heap allocation for database context state management.
		// 2. Concurrency: Shared across all goroutines. Must rely on internal connection pooling to avoid data races.
		// 3. Layout: Struct wrapping the underlying cluster client pointer (`*clientv3.Client`).
		// 4. Failure: If driver configuration crashes, it halts startup; runtime database drops cause handled errors.

		HandlerPod := &handlers.PodHnalder{Store: etcdStore}
		HandlerNode := &handlers.NodeHandler{Store: etcdStore}
		// LIFECYCLE: Core orchestration layer controller handler wrapper.
		// MEMORY: Heap. The pointer address (`&`) escapes the stack because it's bound to the router endpoints.
		// FLOW: Injected into specific HTTP routes to bridge incoming requests directly to the etcd backend store.
		//
		// 1. Allocation: Heap allocation to maintain references to backend persistence layers.
		// 2. Concurrency: Single controller instance; methods (`Get`/`Create`) run concurrently on HTTP goroutines.
		// 3. Layout: Struct memory block holding an interface descriptor ($2$ words) mapping to `etcdStore`.
		// 4. Failure: Internal panics inside handler methods are caught by net/http recovery; server stays alive.

		api.Route("/", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(map[string]string{
					"message": "api Server is running",
				})
			})
		})

		// pods
		api.Post("/namespace/{namespace}/pods/{name}", HandlerPod.Get)
		api.Post("/{namespace}/{name}", HandlerPod.Create)

		// nodes
		api.Get("/namespace/{namespace}/nodes/{name}", HandlerNode.GetNodeHandler)
		api.Post("/{namespace}/{name}", HandlerNode.RegisterNodeHandler)
	})

	return router
}
