package noderesourece

import (
	v1 "k8s.io/api/core/v1"

	"github.com/tekluabayneh/gok8s/types"
	Types "github.com/tekluabayneh/gok8s/types"
)

//
// func ToNodeInfo(n corev1.Node) Types.NodeInfo {
// 	// pull only what the evaluator needs out of the big struct
// 	// hint: n.Status.Allocatable["cpu"] is a resource.Quantity,
// 	// use .MilliValue() for cpu and .Value() for memory
// 	return Types.NodeInfo{
// 	Node *corev1.Node
//
// 	// computed from the node
// 	AllocCPU    int64 // milli
// 	AllocMemory int64 // bytes
// 	AllocPods   int64
//
// 	// maintained by your cache from pod events
// 	Pods         []*corev1.Pod
// 	RequestedCPU int64
// 	RequestedMem int64
// 	UsedPorts    map[int32]bool
//
// }
//
//
// 	}
// }

// check if node fits the pod spec
// check if node healthy  Done
// Unschedulable flag. node.Spec.Unschedulable for cordoned nodes. done
// Pod count. A node has a max number of pods, and Allocatable includes a pods entry.
// nodeSelector. If the pod says nodeSelector: disk=ssd, the node’s labels must match all of them.
// Node affinity. The fancier version of nodeSelector (required rules only; preferred rules belong in scoring).
// Already bound. Skip pods where Spec.NodeName is already set. Only schedule pending, unassigned ones.
// Pod already requested a specific node. If Spec.NodeName is set by the user, it bypasses scheduling entirely.
// Other resources. Ephemeral storage, and GPUs if you ever care.
// Host ports. Two pods can’t both claim the same hostPort on one node.

// ============================================================
// 1. BEFORE SCHEDULING (queue / informer filter, never reaches Filter)
// ============================================================

// Already bound: skip pods where Spec.NodeName is already set.
//   Only pending, unassigned pods go into the scheduling queue.
// Pod pinned to a node by the user: same check as above.
//   Spec.NodeName set means scheduling is bypassed, the kubelet just runs it.

// ============================================================
// 2. FILTER (hard yes/no per node, a node that fails is out)
// ============================================================

// Node health / pressure conditions: Done
//   Ready false, MemoryPressure, DiskPressure, PIDPressure. done
//   (Real k8s turns these into taints; a plain condition check is fine for you.)
// Unschedulable flag: node.Spec.Unschedulable for cordoned nodes. Done

// Resource fit: pod requests <= allocatable - already requested
//    CPU, memory, ephemeral storage, extended resources / GPUs.
// Pod count: Requested pod count + 1 <= Allocatable["pods"].

// nodeSelector: every label in the pod's nodeSelector must match the node's labels.
// Node affinity (required rules only): requiredDuringSchedulingIgnoredDuringExecution.
// Host ports: no two pods can claim the same hostPort on one node.

type Filterresouce struct {
	// Resourcecheker               func() bool
	// NodeSelector                 func() bool
	// Fit                          func() bool
	// Nodehealthpressureconditions func(node *v1.Node) bool
	// Filter                       func(nodes *Types.NodeInfo, p v1.Pod, fesiable []*Types.NodeInfo) ([]*Types.NodeInfo, map[string]string)
	// Node                         func(node *Types.NodeInfo) *v1.Node
	//
}

func NewFilterResource() *Filterresouce {
	return &Filterresouce{}
}

// filter function shouln't now only return bollean value or it couod update cached node or other stuff or return the perfect node to scorer
func (r *Filterresouce) Filter(ni *Types.NodeInfo, p v1.Pod) (*Types.NodeInfo, map[string]string) { // p types need to be change to the exact pod we get way mey need to fetch node data to get ti's spec
	// fesiable := make([]*Types.NodeInfo, 0, len(nodes))
	reson := map[string]string{}
	var fesiable *Types.NodeInfo
	if !r.Nodehealthpressureconditions(ni.Node) {
		reson[ni.Node.Name] = "node is not healthy" // reson need to be update
		return fesiable, reson                      // which mean if the node is not sutable we just return that node cache we get form params only add the node node to fesiable if the node is fesiable node
	}

	// not only return false it may have to be change to something like node2{fit:false, reson:res_alocated } and other stuff
	if !r.NodeSelector(ni.Node, p) {
		reson[ni.Node.Name] = "node NodeSelector is not happy about it" // reson need to be update
		return fesiable, reson
	}

	if !r.Resourcecheker(*ni.Node, p) {
		reson[ni.Node.Name] = "node does not fullfill pods requests resources" // reson need to be update
		return fesiable, reson
	}

	if !r.Fit(p, ni) {
		reson[ni.Node.Name] = "pod does not fit" // reson need to be update
		return fesiable, reson
	}

	fesiable = buildNodeInfo(*ni, p)
	reson[fesiable.Node.Name] = "nod is fesiable" // reson need to be update
	return fesiable, reson
}

func (r *Filterresouce) Nodehealthpressureconditions(n *v1.Node) bool {
	// cfg := spew.ConfigState{Indent: "  ", DisableMethods: true}
	// cfg.Dump(n.Status.Conditions)

	if n.Spec.Unschedulable {
		return false
	}
	for _, C := range n.Status.Conditions {
		switch C.Type {
		case v1.NodeReady:
			if C.Status != v1.ConditionTrue {
				return false
			}

			// this condition must be == not !=  cuz status return negative for the tru reson like if it has it will say false
		case v1.NodeMemoryPressure, v1.NodeDiskPressure, v1.NodePIDPressure:
			if C.Status == v1.ConditionTrue {
				return false
			}
		}
	}
	return true
}

func (r *Filterresouce) Resourcecheker(n v1.Node, p v1.Pod) bool {
	// n.Status.NodeInfo.OSImage // need to check if pods fits with it like if pod may need to run i certain criterrica
	//
	label := n.Labels
	// HOT: this failteration may need to be remove as it prevent all pod to be assing in the control-plane node but static pod can be assined here and may need to use n.Spec.Taint
	for leb := range label {
		if leb == "node-role.kubernetes.io/control-plane" {
			return false
		}
	}
	// node-role.kubernetes.io/control-plane"
	// getNodeRoles(item)?.length > 0 ? getNodeRoles(item) : "<none>",

	return true
}

func (r *Filterresouce) NodeSelector(n *v1.Node, p v1.Pod) bool {
	return true
}

// Resource fit: pod requests <= allocatable - already requested
//    CPU, memory, ephemeral storage, extended resources / GPUs.
// Pod count: Requested pod count + 1 <= Allocatable["pods"].

func (r *Filterresouce) Fit(p v1.Pod, nodeinfo *Types.NodeInfo) bool {
	if nodeinfo.Node.Status.Allocatable.Cpu().MilliValue() <= p.Status.AllocatedResources.Cpu().MilliValue() {
		return false
	}

	if nodeinfo.Node.Status.Allocatable.Memory().MilliValue() <= p.Status.AllocatedResources.Memory().MilliValue() {
		return false
	}
	// TODO get the cached pods count form the cache and compare it with limted one
	if int64(len(nodeinfo.Pods)) >= nodeinfo.Node.Status.Allocatable.Pods().Value() {
		return false
	}
	return true
}

func buildNodeInfo(node Types.NodeInfo, po v1.Pod) *Types.NodeInfo {
	alloc := node.Node.Status.Allocatable
	no := &types.NodeInfo{
		Node:        node.Node,
		AllocCPU:    alloc.Cpu().MilliValue(),
		AllocMemory: alloc.Memory().Value(),
		AllocPods:   alloc.Pods().Value(),
		UsedPorts:   map[int32]bool{},
		// RequestedCPU: po.Status.Resources.Requests.Cpu().Value(), // this tow are total pods cpu value
		// 	RequestedMem: po.Status.Resources.Requests.Memory().MilliValue(), this one is also same total of memory value for pods in the patricual node
	}

	return no
}
