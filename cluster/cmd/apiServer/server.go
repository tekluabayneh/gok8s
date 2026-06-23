package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tekluabayneh/gok8s/cmd/apiServer/routers"
)

type APIServerType struct {
	router http.Handler

	// 1. Stack or heap?
	// Depends on instantiation. Since an API server runs forever, this will escape to the heap.

	// 2. Passing value, pointer, or reference—and who owns it?
	// Value (interface), but contains underlying pointers. The router creator owns data; this struct shares it.

	// 3. Can multiple threads/goroutines touch this at the same time?
	// Yes. Concurrent reads across HTTP goroutines are 100% thread-safe; writes at runtime would race.

	// 4. What is the source of truth—and where is the state machine?
	// Source of truth is the concrete router (e.g., Chi). State machine is its internal URL matching tree.

	// 5. When this loop/function panics or errors out, what state is left behind?
	// net/http recovers panics per request. This router instance is left completely untouched and intact.

	// 6. What is the absolute worst-case time and space complexity here?
	// Space: O(N) for N total routes. Time: O(K) where K is the length of the incoming URL path string.

	// 7. What system call does this trigger under the hood?
	// None for the struct itself. Serving requests triggers network I/O syscalls like read, write, or epoll.

	// 8. What happens to the file descriptor, socket, or pipe if this process dies?
	// The OS kernel immediately reclaims memory and forcibly closes all open network sockets.

	// 9. Why this specific data structure shape and not another?
	// Using the http.Handler interface lets you decouple and swap any router without changing server code.

	// 10. What happens if the network latency spikes or a timeout occurs?
	// Struct is fine, but handlers leak/waste resources unless you explicitly monitor req.Context().Done().
}

func AppAPIServerNew() *APIServerType {
	return &APIServerType{
		router: routers.LoadRouter(),
	}
	// 1. Stack or heap?
	// Heap. You are returning a pointer (`*APIServerType`), forcing the struct to escape the function stack.

	// 2. Passing value, pointer, or reference—and who owns it?
	// Returning a pointer value. The calling function takes ownership of this specific server instance pointer.

	// 3. Can multiple threads/goroutines touch this at the same time?
	// Yes. Goroutines can safely read the router concurrently. Instantiation itself happens on a single thread.

	// 4. What is the source of truth—and where is the state machine?
	// Source of truth is the initialized `routers.LoadRouter()` return value allocated inside this constructor.

	// 5. When this loop/function panics or errors out, what state is left behind?
	// If this initialization panics, the server never starts, and memory is garbage collected immediately.

	// 6. What is the absolute worst-case time and space complexity here?
	// Space: O(N) to store the route tree in memory. Time: O(1) allocation for the wrapper struct itself.

	// 7. What system call does this trigger under the hood?
	// None directly. It only executes user-space memory allocations on the heap via the Go runtime.

	// 8. What happens to the file descriptor, socket, or pipe if this process dies?
	// If the process dies here, the OS kernel instantly cleans up and closes any un-inherited open descriptors.

	// 9. Why this specific data structure shape and not another?
	// Clean factory pattern encapsulation. It bundles the router dependencies neatly before mounting to the server.

	// 10. What happens if the network latency spikes or a timeout occurs?
	// Initialization doesn't care about network spikes; it only affects active request streams later on.
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

	// LIFECYCLE: it's long lived LIFECYCLE, it initialization server with nessory methond like port injecting router handler and it's with in APIServerStart() LIFECYCLE when the function finishe it will but since this is long lived and it live in the heap the server run forever so garbage wont' touch it and it's http server pointer value referance
	// MEMORY: Escapes to the Heap because a pointer is returned. GC tracks it until the server shuts down.
	// FLOW: Called once at startup on the main thread; the it sits on the proccessore shared globally.
	//
	// 1. Allocation: Heap allocation due long lived session
	// 2. Concurrency: it will always be single thread, but multiple goroutines can run inside it or based on request
	// 3. Layout: Struct initialization with the aproprate server methods
	// 4. Failure: if server initialization has problem it will panic and stop the server so os kernel will wipe out from process
	// 5. Cost: O(1) space/time for the struct wrapper wrapper; runtime tax is a single heap allocation.
	//

	fmt.Printf("Server is running on post %v", PORT)
	if err := apiServ.ListenAndServe(); err != nil {
		log.Fatal("Server failed to run ", err.Error())
	}
}
