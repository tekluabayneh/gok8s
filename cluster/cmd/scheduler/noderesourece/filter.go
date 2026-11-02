package noderesourece

import (
	v1 "k8s.io/api/core/v1"

	"github.com/tekluabayneh/gok8s/types"
	Types "github.com/tekluabayneh/gok8s/types"
)

type Filterresouce struct{}

func NewFilterResource() *Filterresouce {
	return &Filterresouce{}
}

// filter function shouln't now only return bollean value or it couod update cached node or other stuff or return the perfect node to scorer
func (r *Filterresouce) Filter(ni *Types.NodeInfo, p v1.Pod) (*Types.NodeInfo, map[string]string) { // p types need to be change to the exact pod we get way mey need to fetch node data to get ti's spec
	// fesiable := make([]*Types.NodeInfo, 0, len(nodes))
	reson := map[string]string{}
	var fesiable *Types.NodeInfo
	if !r.Nodehealthpressureconditions(ni.Node) {
		reson[ni.Node.Name] = "node is not healthy"
		return fesiable, reson // which mean if the node is not sutable we just return that node cache we get form params only add the node node to fesiable if the node is fesiable node
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
	label := n.Labels
	for leb := range label {
		if leb == "node-role.kubernetes.io/control-plane" {
			return false
		}
	}

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
		Pods:        []*v1.Pod{},
		AllocPods:   alloc.Pods().Value(),
		UsedPorts:   map[int32]bool{},
		// RequestedCPU: po.Status.Resources.Requests.Cpu().Value(), // this tow are total pods cpu value
		// 	RequestedMem: po.Status.Resources.Requests.Memory().MilliValue(), this one is also same total of memory value for pods in the patricual node
	}

	return no
}
