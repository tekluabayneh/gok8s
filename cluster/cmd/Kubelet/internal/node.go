package internal

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	goruntime "runtime"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/tekluabayneh/gok8s/types"
	"github.com/tekluabayneh/gok8s/utils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Kub struct {
	types.Kubelet
}

var (
	defaultbool bool
	HostName    = "kind-control-planemm"
)

func NewKubelet(kl *Kub) *types.Kubelet {
	newNode, err := kl.RegisterNode()
	if err != nil {
		panic(err)
	}

	// cfg := spew.ConfigState{
	// 	Indent:                  "  ",
	// 	DisableMethods:          true,
	// 	DisablePointerAddresses: true,
	// 	DisableCapacities:       true,
	// 	SortKeys:                true,
	// 	SpewKeys:                true,
	// 	MaxDepth:                4,
	// }
	//
	b, err := json.MarshalIndent(newNode, "", "")
	if err != nil {
		panic(nil)
	}

	json := utils.YamlToJson(string(b))
	fmt.Println(string(json))
	// cfg.Dump(json)

	newkubelet := &types.Kubelet{
		// nodeName: newNode.Name,
	}

	return newkubelet
}

func (kl *Kub) initialNode() (v1.Node, error) {
	node := v1.Node{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Node",
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"node.alpha.kubernetes.io/ttl":                           "0",
				"volumes.kubernetes.io/controller-managed-attach-detach": "true",
			},
			Labels: map[string]string{
				v1.LabelHostname:   HostName,
				v1.LabelOSStable:   goruntime.GOOS,
				v1.LabelArchStable: goruntime.GOARCH,
				// kubeletapis.LabelOS:   goruntime.GOOS,
				// kubeletapis.LabelArch: goruntime.GOARCH,
				//
			},
		},
	}
	return node, nil
}

// FLAGS used below:
//   [DAEMON]       background loop or watch, the kubelet talking TO the API server
//   [HTTPS SERVER] endpoint the API server calls ON the kubelet (none in Tier 1)

// ============================================================
// 1. NODE REGISTRATION, STATUS AND HEARTBEAT   [DAEMON]
// ============================================================

// TODO: registerNode
// Goal: when the kubelet starts, make sure a Node object for this machine
// exists on the API server, so the scheduler can see it.
// Steps: build a v1.Node (name, labels incl. kubernetes.io/hostname),
// fill capacity and allocatable, then Create it, or Update if it already exists.
func (k *Kub) RegisterNode() (*v1.Node, error) {
	node, err := k.initialNode()
	if err != nil {
		panic(err)
	}
	var ip string

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	adress, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, add := range adress {
		if ipn, ok := add.(*net.IPNet); ok && !ipn.IP.IsLoopback() && ipn.IP.To4() != nil {
			ip = ipn.IP.String()
			break
		}
	}

	node.Name = HostName      // and this name also must be some kida programatic not static typed like this
	node.ResourceVersion = "" // same here

	node.Status = v1.NodeStatus{
		Addresses: []v1.NodeAddress{
			{
				Type:    v1.NodeInternalIP,
				Address: ip,
			},
			{
				Type:    v1.NodeHostName,
				Address: hostname,
			},
		},
	}
	node.Status.DaemonEndpoints = v1.NodeDaemonEndpoints{
		KubeletEndpoint: v1.DaemonEndpoint{
			Port: 10250,
		},
	}
	node.Status.Features = &v1.NodeFeatures{
		SupplementalGroupsPolicy: &defaultbool,
	}

	node.Status.Images = []v1.ContainerImage{
		// {
		// 	Names:     []string{""},
		// 	SizeBytes: 24325, // this SizeBytes is just random so it aslo need to be change
		// },
		//
	}
	bootId, err := os.ReadFile("/proc/sys/kernel/random/boot_id")
	if err != nil {
		panic(err)
	}

	SystemUUID, err := os.ReadFile("/sys/class/dmi/id/product_uuid")
	if err != nil {
		panic(err)
	}

	KernelVersion, err := host.KernelVersion()
	if err != nil {
		panic(err)
	}

	node.Status.NodeInfo = v1.NodeSystemInfo{
		Architecture:            runtime.GOARCH,
		BootID:                  strings.TrimSpace(string(bootId)),
		ContainerRuntimeVersion: "containerd://2.2.0",
		KernelVersion:           KernelVersion,
		KubeProxyVersion:        "v1.0.0",
		KubeletVersion:          "v1.35.0",
		MachineID:               utils.UuidGenerator(),
		OperatingSystem:         runtime.GOOS,
		OSImage:                 "",
		Swap:                    &v1.NodeSwapStatus{
			// Capacity: &int64(3986681856),
		},
		SystemUUID: strings.TrimSpace(string(SystemUUID)),
	}

	node.Status.RuntimeHandlers = []v1.NodeRuntimeHandler{
		{
			Name: "", // name here need to be udpated this has to be came from the image runtime so i will fill up later on when i implement the runtime
			Features: &v1.NodeRuntimeHandlerFeatures{
				RecursiveReadOnlyMounts: &defaultbool,
				UserNamespaces:          &defaultbool,
			},
		},
	}

	// node.nore nodeResources
	NodeCapacity, NodeAllocatable := k.NodeResources()

	node.Status.Allocatable = NodeAllocatable
	node.Status.Capacity = NodeCapacity

	return &node, nil
}

// TODO: nodeResources
// Goal: report how much CPU and memory this machine has.
// This is the data scoreMostAllocated reads from Allocatable.
// Steps: CPU = cores * 1000 (millicores), memory = MemTotal from /proc/meminfo
// converted to bytes, allocatable = capacity minus a small system reserve.
func (kl *Kub) NodeResources() (capacity, allocatable v1.ResourceList) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	NodeAllocatable := v1.ResourceList{
		v1.ResourceCPU:              resource.MustParse(fmt.Sprintf("%d", runtime.NumCPU())),
		v1.ResourceMemory:           resource.MustParse(fmt.Sprintf("%d", memory.Total)),
		v1.ResourceEphemeralStorage: resource.MustParse("10Gi"),
		v1.ResourcePods:             resource.MustParse("110"),

		v1.ResourceName(v1.ResourceHugePagesPrefix + "1Gi"): resource.MustParse("0"),
		v1.ResourceName(v1.ResourceHugePagesPrefix + "2Mi"): resource.MustParse("0"),
	}

	NodeCapacity := v1.ResourceList{
		v1.ResourceCPU:              resource.MustParse("4"),
		v1.ResourceMemory:           resource.MustParse("4Gi"),
		v1.ResourceEphemeralStorage: resource.MustParse("10Gi"),
		v1.ResourcePods:             resource.MustParse("110"),

		v1.ResourceName(v1.ResourceHugePagesPrefix + "1Gi"): resource.MustParse("0"),
		v1.ResourceName(v1.ResourceHugePagesPrefix + "2Mi"): resource.MustParse("0"),
	}

	return NodeAllocatable, NodeCapacity
}

// TODO: syncNodeStatus
// Goal: tell the API server what state the node is in (Ready, pressure conditions).
// Steps: build the conditions, update Node.Status, run on a timer.
// Real kubelet: sends on change, and otherwise only every few minutes. Check the source.
func (kl *Kub) SyncNodeStatus() error {
	node, err := kl.initialNode()
	if err != nil {
		panic(err)
	}
	node.Status.Conditions = []v1.NodeCondition{ // we may now don't need to add this condtion since condition are based on how that particular kubelet operate and ongiong opration
		{
			// LastTransitionTime: "2026-09-04T14:07:28Z",
			Message: "kubelet has sufficient memory available",
			Reason:  "KubeletHasSufficientMemory",
			Status:  "True",
			Type:    "Ready-fake",
		},
	}

	return nil
}

// TODO: renewLease
// Goal: the cheap heartbeat. Update one small Lease object (named after the node,
// in the kube-node-lease namespace) every ~10 seconds, so the control plane
// can tell the node is alive without receiving the whole Node status each time.
// func (k *Kubelet) renewLease() error { return nil }

// ============================================================
// 2. POD CONFIG SOURCE   [DAEMON]
// ============================================================

// TODO: startPodSource
// Goal: learn about every pod that is assigned to THIS node, and nothing else.
// Steps:
//   1. Start your informer on pods with a field selector spec.nodeName == this node.
//   2. On add, update, delete events, push a small update message
//      (type + pod) into a channel that the pod workers read from.
//   3. Use the pod UID as the identity, not the name, because a pod can be
//      deleted and recreated with the same name.
// func (k *Kubelet) startPodSource(updates chan<- PodUpdate) {}

// ============================================================
// 3. RUNTIME MANAGER AND IMAGE MANAGER   [DAEMON]
// ============================================================

// TODO: define a Runtime interface
// Goal: the thin layer between the kubelet and the real runtime (containerd, Docker),
// so the rest of the kubelet never calls the runtime directly.
// Methods to include: PullImage, CreateContainer, StartContainer,
// StopContainer, ListContainers, ContainerStatus.
// Write a fake in-memory version first, so you can test the sync loop without a runtime.
// type Runtime interface{}

// TODO: ensureImage
// Goal: pull an image only when needed.
// Steps: check imagePullPolicy (Always, IfNotPresent, Never) against whether
// the image already exists on the node, then pull if the policy says so.
// Return an error on failure, since the sync loop uses it to report ErrImagePull.
// func (k *Kubelet) ensureImage(container v1.Container) error { return nil }

// ============================================================
// 4. POD WORKERS AND SYNC   [DAEMON]
// ============================================================

// TODO: dispatchPodUpdate
// Goal: give each pod its own worker goroutine, so updates to ONE pod run
// in order, but different pods don't block each other.
// Steps: keep a map UID -> channel, start a worker the first time a UID
// is seen, and send the update to that pod's worker.
// func (k *Kubelet) dispatchPodUpdate(u PodUpdate) {}

// TODO: syncPod
// Goal: make reality match the pod spec. This is the core of the kubelet.
// Steps:
//   1. If the pod has a DeletionTimestamp, stop its containers and finish.
//   2. Ask the runtime what is actually running for this pod.
//   3. Compare with pod.Spec.Containers: start missing containers
//      (ensureImage first), stop ones that shouldn't exist.
//   4. Hand the result to the status manager.
// Must be safe to run repeatedly: running it twice must not start duplicates.
// func (k *Kubelet) syncPod(pod *v1.Pod) error { return nil }

// ============================================================
// 5. STATUS MANAGER   [DAEMON]
// ============================================================

// TODO: setPodStatus / syncStatus
// Goal: write the pod's phase and container statuses back to the API server.
// Steps:
//   1. Keep the latest status per pod UID in a cache.
//   2. Compute the phase: Pending (not all started), Running (at least one running),
//      Succeeded / Failed (all finished, depending on restartPolicy).
//   3. Only send an update when the status changed, to avoid hammering the API server.
//   4. Write to the pod's status, not the pod spec.
// func (k *Kubelet) setPodStatus(pod *v1.Pod, statuses []ContainerState) {}

// ============================================================
// 6. EVENT RECORDER   [DAEMON]
// ============================================================

// TODO: recordEvent
// Goal: leave a trail of what happened, so `describe pod` shows events.
// Events to emit from syncPod: Pulling, Pulled, Created, Started, Failed, BackOff.
// Rules: events are best-effort. Never let a failed event block or fail the sync.
// func (k *Kubelet) recordEvent(pod *v1.Pod, reason, message string) {}
