package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/tekluabayneh/gok8s/internals/scheduler"
	Types "github.com/tekluabayneh/gok8s/types"
)

type cacheImpl struct {
	stop   <-chan struct{}
	period time.Duration

	// This mutex guards all fields within this cache struct.
	mu sync.RWMutex

	// podStates map[string]*podState
	nodes map[string]*Types.NodeInfo
}

func main() {
	// Sigctx, stop := signal.NotifyContext(
	// 	context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer stop()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	w := scheduler.NewWatcher("http://localhost:8080/api/v1/pods") // HOT: this hard coded url need to be changed
	s := scheduler.New()
	event := make(chan Types.WatchEvent)
	go s.Run(ctx, event)

	err := w.Watch(ctx, event) // this function is blocking but if the select case return it will close and return value or nil or error
	close(event)

	if err != nil && ctx.Err() == nil {
		log.Println("watch stopped:", err)
	}

	fmt.Println("couse=>", context.Cause(ctx))
}

// =====================================================================
// goK8s SCHEDULER: TODO + READING GUIDE
// Rule: read the real source, then write YOUR OWN logic. Comments only here.
// Real k8s paths can shift between versions, so grep if a path is missing.
// =====================================================================

// ---------------------------------------------------------------------
// PHASE 1: CONFIG (flags + kubeconfig)
// ---------------------------------------------------------------------
// TODO: parse flags: --kubeconfig, --master (optional API url override),
//       --healthz-port, --leader-elect, --v (log level)
// TODO: kubeconfig fallback order: flag -> $KUBECONFIG -> ~/.kube/config
//       (reuse your CLI kubeconfig + yaml-to-json code, don't rewrite it)
// TODO: from kubeconfig extract: server URL, CA cert, token OR client cert/key
// TODO: validate at startup and fail fast with a clear error if missing/unreachable
// TODO: optional later: in-cluster mode (env vars + service account token file)
//
// READ:
//   client-go/tools/clientcmd/loader.go      (how kubeconfig is located/merged)
//   client-go/rest/config.go                 (rest.Config, InClusterConfig)
//   cmd/kube-scheduler/app/options/options.go (scheduler flags and config loading)
//   cmd/kube-scheduler/app/server.go          (Run(): startup order of everything)

// ---------------------------------------------------------------------
// PHASE 2: PROCESS LIFECYCLE (signals, shutdown, daemon)
// ---------------------------------------------------------------------
// TODO: create a root context; cancel it on SIGINT/SIGTERM
// TODO: pass ctx into informers, queue, and the scheduling loop
// TODO: on shutdown: stop informers, finish the in-flight bind, then exit
// TODO: log to stdout/stderr only (no daemonizing inside Go)
// TODO: write systemd unit: ExecStart, Restart=always, User, After=network.target
// TODO: quick test without systemd: nohup ./scheduler ... > sched.log 2>&1 &
//
// READ:
//   Go docs: os/signal NotifyContext
//   man systemd.service (ExecStart, Restart, Type=simple)
//   man systemd.unit (After, Wants)

// ---------------------------------------------------------------------
// PHASE 3: API CALLS (the routes your scheduler needs)
// ---------------------------------------------------------------------
// TODO: LIST + WATCH pods   -> GET /api/v1/pods, then ?watch=true
//       (only queue pods where spec.nodeName is empty)
// TODO: LIST + WATCH nodes  -> GET /api/v1/nodes, then ?watch=true
// TODO: BIND a pod          -> POST /api/v1/namespaces/{ns}/pods/{name}/binding
// TODO: mark unschedulable  -> PATCH .../pods/{name}/status
//       (PodScheduled=False, reason Unschedulable)
// TODO: emit events         -> POST /api/v1/namespaces/{ns}/events
// TODO: handle 409 on bind (pod already bound or deleted) and 404
// TODO: your API server must implement the binding route, check that first
//
// READ:
//   pkg/registry/core/pod/storage/storage.go  (BindingREST: what bind does server-side)
//   staging/src/k8s.io/api/core/v1/types.go   (Binding, PodCondition structs)

// ---------------------------------------------------------------------
// PHASE 4: INFORMERS + CACHE
// ---------------------------------------------------------------------
// TODO: pod informer built on YOUR Reflector + DeltaFIFO
// TODO: node informer, same way
// TODO: scheduler cache: nodes + their allocated resources (sum of pods per node)
// TODO: wait for caches to sync before the loop starts
// TODO: test: kill the API server mid-run, confirm watch reconnect + relist works
//
// READ:
//   pkg/scheduler/eventhandlers.go            (what events update cache/queue)
//   pkg/scheduler/backend/cache/cache.go      (older: pkg/scheduler/internal/cache)

// ---------------------------------------------------------------------
// PHASE 5: SCHEDULING QUEUE
// ---------------------------------------------------------------------
// TODO: queue of unscheduled pods (start with a simple FIFO)
// TODO: unschedulable set + retry with backoff when no node fits
// TODO: re-queue a pod when a node is added or a pod is deleted
//
// READ:
//   pkg/scheduler/backend/queue/scheduling_queue.go
//   (older: pkg/scheduler/internal/queue/scheduling_queue.go)

// ---------------------------------------------------------------------
// PHASE 6: THE LOOP (filter -> score -> pick -> bind)
// ---------------------------------------------------------------------
// TODO: pop pod -> filter nodes -> score nodes -> pick best -> bind
// TODO: filter v1: enough CPU/memory (requests vs node allocatable minus used)
// TODO: filter v2: nodeSelector, taints/tolerations, nodeName
// TODO: score v1: least allocated (most free resources wins)
// TODO: assume the pod on the cache BEFORE the bind call finishes,
//       forget it if the bind fails
// TODO: bind in a goroutine so the loop keeps going
//
// READ:
//   pkg/scheduler/schedule_one.go             (scheduleOne, assume, bind)
//   pkg/scheduler/framework/interface.go      (Filter / Score plugin shapes)
//   pkg/scheduler/framework/plugins/noderesources/fit.go
//   pkg/scheduler/framework/plugins/noderesources/least_allocated.go

// ---------------------------------------------------------------------
// PHASE 7: HARDENING
// ---------------------------------------------------------------------
// TODO: /healthz and /readyz HTTP endpoints
// TODO: structured logs with levels
// TODO: keep NO state on disk; rebuild everything from list on restart
// TODO: leader election (Lease in coordination.k8s.io) if running 2+ copies
// TODO: retry on startup if the API server is down
//
// READ:
//   client-go/tools/leaderelection/leaderelection.go
//   client-go/tools/leaderelection/resourcelock/leaselock.go

// ---------------------------------------------------------------------
// FIRST MILESTONE (do this before anything else)
// ---------------------------------------------------------------------
// TODO: scheduler loads kubeconfig, connects to your API server, and just
//       PRINTS every unscheduled pod it sees. Nothing else.
