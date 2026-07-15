etcd                → pure server (responds to reads/writes)
api-server          → pure server (responds to HTTP requests)

scheduler           → pure daemon (watches and assigns pods)
controller-manager  → pure daemon (watches and reconciles state)

<!-- this tow belong to the node worker -->
kubelet             → both (daemon + server)
    <!---->
    <!-- goroutines -->
    <!-- channels — for communication between components -->
    <!-- select statement — kubelet has multiple things happening at once -->
    <!-- sync.WaitGroup — waiting for multiple goroutines -->
    <!-- sync.Mutex — protecting shared state like pod list -->
    <!-- context — cancellation, timeouts, shutting down gracefully -->a

      <!--     Node self-registration — on startup, kubelet creates/updates its own Node object in the API server via the apiserver's REST API (not etcd directly — kubelet never talks to etcd directly, only through apiserver). Reports capacity, allocatable resources, conditions. -->

      <!-- Pod source aggregation (PodConfig) — merges pod specs from three sources into one stream: apiserver (watch), static pod manifest files on disk, and an HTTP endpoint. This is a config multiplexer, not just one watch. -->

      <!-- syncLoop — the central event loop. It's a select over multiple channels: pod config changes, a periodic sync timer, PLEG events, and probe manager events. Every event funnels through here and gets dispatched to syncLoopIteration. -->

      <!-- PLEG (Pod Lifecycle Event Generator) — this is the piece most tutorials skip and it's genuinely important for your project. Instead of syncLoop polling the container runtime directly for every pod on every tick (expensive at scale), PLEG periodically relists containers via CRI, diffs against its own cache, and emits discrete events (ContainerStarted, ContainerDied, etc.) onto a channel that syncLoop consumes. Newer "Evented PLEG" replaces polling with gRPC streaming events from the CRI runtime directly — worth understanding both since it's a nice real example of polling-vs-event-driven tradeoffs, which is squarely a systems-design question you already like. -->

      <!-- PodWorkers — one worker (goroutine) per pod, serializing all sync operations for that specific pod so you never race two syncs of the same pod against each other. This is a very learnable concurrency pattern for goK8s. -->

      <!-- CRI (Container Runtime Interface) client — kubelet is a gRPC client to the container runtime (containerd/CRI-O). It never touches containers directly. Create/start/stop/remove sandbox and containers, pull images, all via CRI gRPC calls. -->

      <!-- CNI — invoked by the container runtime (or kubelet, depending on version) to set up pod networking. -->

      <!-- Probes — liveness/readiness/startup probes run on their own timers, feed results into syncLoop via ProbeManager's channel. -->

      <!-- Status reporting — periodically pushes pod status and node status back to the apiserver (not the same channel as receiving config — this is outbound). -->

      <!-- cAdvisor — embedded, scrapes node + container resource usage, feeds eviction manager and exposes /metrics for the metrics-server. -->

      <!-- Eviction manager — watches node resource pressure (memory, disk), evicts pods gracefully before the kernel OOM-killer has to step in — this is a good example of "the two very different pod-death paths" that trips people up. -->

      <!---->

kube-proxy          → pure daemon (manages network rules)

// <https://claude.ai/chat/8213157e-1d17-4ede-9c71-e4c51dee6cba>
//<https://gemini.google.com/app/5b421043e5cfc468>
