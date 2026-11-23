package types

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

type NodeInfo struct {
	Node *corev1.Node

	// computed from the node
	AllocCPU     int64
	AllocMemory  int64
	AllocPods    int64
	Pods         []*corev1.Pod
	RequestedCPU int64
	RequestedMem int64
	UsedPorts    map[int32]bool
}

//
// type PodInfo struct {
// 	Pod *v1.Pod
//
// 	// pre-parsed affinity terms, so Filter/Score plugins don't re-parse the spec
// 	RequiredAffinityTerms      []AffinityTerm
// 	RequiredAntiAffinityTerms  []AffinityTerm
// 	PreferredAffinityTerms     []WeightedAffinityTerm
// 	PreferredAntiAffinityTerms []WeightedAffinityTerm
// }

type (
	EventType  string
	WatchEvent struct {
		Type   EventType `json:"type"`
		Object v1.Pod    `json:"object"`
	}
)
