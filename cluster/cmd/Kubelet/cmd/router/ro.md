Kubelet's ports and REST API
Port 10250 (main — HTTPS, authenticated)

Default secure port. The apiserver uses TLS client certificates to talk to it.
Most endpoints require authentication (bearer token or client cert).
Root / returns 404 by default.

Port 10248 (healthz — HTTP, unauthenticated)

Simple HTTP health check endpoint on localhost only.
/healthz returns 200 OK if kubelet is alive.
This is for kubelet to prove it's running, not for full control.

Port 10255 (read-only port — deprecated but still used)

Some setups expose pod metadata without auth on localhost.
Returns pod lists, basic metrics. Disabled in newer clusters by default.

Port 4194 (cAdvisor metrics)

Embedded cAdvisor monitoring interface.
Node and container resource usage.

What endpoints are on port 10250
The kubelet REST API is not officially documented in the Kubernetes docs (which is frustrating), but these are the real endpoints:

/pods — GET: List of all pods on this node (what kubelet sees)
/logs/{pod}/{container} — GET: Stream container logs (what kubectl logs uses)
/exec/{pod}/{container} — POST: Execute command in container with stdin/stdout/stderr (what kubectl exec uses, bidirectional over WebSocket/SPDY)
/metrics — GET: Prometheus-format node + pod metrics
/stats — GET: Node and container statistics (used by metrics-server)
/spec — GET: Node hardware spec
/proxy/{path} — Proxy requests to other kubelet endpoints

How apiserver uses it
The apiserver:

Calls /pods periodically to see what's actually running vs. desired state
Calls /logs when you run kubectl logs podname
Calls /exec when you run kubectl exec podname -- bash
Calls /metrics to get node capacity and usage (for scheduler decisions)
Never directly tells kubelet "run this pod" — instead, it updates the Pod object in etcd, and kubelet's syncLoop watches for the change
