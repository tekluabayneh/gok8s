✅ Pod CRUD + etcd store ← you're here
Watch mechanism ← do this next, everything depends on it
Node registration (kubelet POSTs to /api/v1/nodes)
Status subresources
Remaining workload resources (Deployment → ReplicaSet → Pod chain)
Service + Endpoints
ConfigMap, Secret
Admission pipeline (stub → real RBAC)
