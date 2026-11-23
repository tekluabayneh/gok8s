package handlers

//
// import (
// 	"fmt"
//
// 	corev1 "k8s.io/api/core/v1"
//
// 	Types "github.com/tekluabayneh/gok8s/types"
// )
//
// type cache struct{}
//
// // func ToNodeInfo(n corev1.Node) Types.NodeInfo {
// // 	// pull only what the evaluator needs out of the big struct
// // 	// hint: n.Status.Allocatable["cpu"] is a resource.Quantity,
// // 	// use .MilliValue() for cpu and .Value() for memory
// // 	return Types.NodeInfo{
// // 		Name:          n.Name,
// // 		Ready:         false, // for now but it has to be update
// // 		Unschedulable: false, // same
// // 		AllocCPU:      34,    // same
// // 		AllocMemory:   24,    // same
// // 		Taints:        []corev1.Taint{},
// // 	}
// // }
//
// // check if node fits the pod spec
// // check if node healthy  Done
// // Unschedulable flag. node.Spec.Unschedulable for cordoned nodes. done
// // Pod count. A node has a max number of pods, and Allocatable includes a pods entry.
// // nodeSelector. If the pod says nodeSelector: disk=ssd, the node’s labels must match all of them.
// // Node affinity. The fancier version of nodeSelector (required rules only; preferred rules belong in scoring).
// // Already bound. Skip pods where Spec.NodeName is already set. Only schedule pending, unassigned ones.
// // Pod already requested a specific node. If Spec.NodeName is set by the user, it bypasses scheduling entirely.
// // Other resources. Ephemeral storage, and GPUs if you ever care.
// // Host ports. Two pods can’t both claim the same hostPort on one node.
// // Pressure conditions. MemoryPressure, DiskPressure and PIDPressure show up in node conditions, and real clusters turn them into taints.
//
// func Filter(n corev1.Node, p Types.WatchEvent) bool { // p types need to be change to the exact pod we get way mey need to fetch node data to get ti's spec
// 	// n.Status.NodeInfo.OSImage // need to check if pods fits with it like if pod may need to run i certain criterrica
//
// 	limitofnode := n.Status.Allocatable.Pods()
// 	fmt.Println("list of pods", limitofnode.Size())
//
// 	// instant false if not is not ready
// 	if n.Status.Conditions[len(n.Status.Conditions)-1].Type != "Ready" || n.Spec.Unschedulable {
// 		return false
// 	}
// 	label := n.Labels
//
// 	if n.Status.Allocatable.Cpu().MilliValue() <= 23 || n.Status.Allocatable.Memory().MilliValue() <= 23 { // the number 23 need to be change they have to be from pod data
// 		return false
// 	}
//
// 	// HOT: this failteration may need to be remove as it prevent all pod to be assing in the control-plane node but static pod can be assined here and may need to use n.Spec.Taint
// 	for leb := range label {
// 		if leb == "node-role.kubernetes.io/control-plane" {
// 			fmt.Println("got it", leb)
// 			return false
//
// 		}
// 	}
// 	// node-role.kubernetes.io/control-plane"
// 	// getNodeRoles(item)?.length > 0 ? getNodeRoles(item) : "<none>",
//
// 	return false
// }
//
// func Score(n Types.NodeInfo, p Types.WatchEvent) int {
// 	return 23
// }
//
// // pattern must be used
// // strategy parretn
// // factory pattern
