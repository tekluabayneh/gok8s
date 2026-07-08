package etcd

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func StartResourceInformer(ctx context.Context, client *clientv3.Client, resType string) {
	fmt.Printf("initialized %s watcher \n", resType)
	watcher := client.Watch(ctx, resType, clientv3.WithPrefix())
	for response := range watcher { // this is blocking wating for new change and grpc will stream to this channel if there is not data scheduler will park this routing
		for _, event := range response.Events {
			/// based on the even time call the Delta's function
			switch event.Type {
			case mvccpb.PUT:
				if event.IsCreate() {
					Fifo.Add(event.Kv.Key, event.Kv.Value, event.Kv.ModRevision)
				} else {
					Fifo.Update(event.Kv.Key, event.Kv.Value, event.Kv.ModRevision)
				}
			case mvccpb.DELETE:
				Fifo.Delete(event.Kv.Key, event.Kv.ModRevision)
			}
		}
	}
}

/*
================================================================================
CORE KUBERNETES INFORMER / WATCHER CHECKLIST
================================================================================

1. Pods ("gok8s/Pod")
   - Handled: YES
   - Purpose: Monitors the state, spec, and phase of scheduled/unscheduled containers.
   - Core consumer: Scheduler (to find nodes) and Kubelet (to run containers).

2. Nodes ("gok8s/Node")
   - Handled: YES
   - Purpose: Tracks physical/virtual machine health, capacities (CPU/RAM), and status.
   - Core consumer: Scheduler (to evaluate capacity) and Controller Manager (evictions).

3. Services ("gok8s/Service")
   - Handled: NO (To do)
   - Purpose: Tracks the stable virtual IP definitions and ports for application access.
   - Core consumer: Kube-Proxy / Networking Layer.

4. Endpoints / EndpointSlices ("gok8s/Endpoints")
   - Handled: NO (To do)
   - Purpose: Dynamically maps a Service to the raw backend IPs of the actual working Pods.
   - Core consumer: Kube-Proxy (to update iptables/IPVS routing tables).

5. Deployments / ReplicaSets ("gok8s/Deployment")
   - Handled: NO (To do)
   - Purpose: Monitors high-level desired application state, replication factors, and rollouts.
   - Core consumer: Deployment/ReplicaSet Controllers (to scale up/down or create Pods).

6. ConfigMaps / Secrets ("gok8s/ConfigMap", "gok8s/Secret")
   - Handled: NO (To do)
   - Purpose: Stores configuration files, environment variables, and authentication tokens.
   - Core consumer: Kubelet (to mount data into containers at runtime).

7. Namespaces ("gok8s/Namespace")
   - Handled: NO (To do)
   - Purpose: Scopes resources and handles cascade deletions (deleting a namespace deletes its pods).
   - Core consumer: API Server and Resource Cleaners.
================================================================================
*/
