package config

//
// type Node struct {
// 	metav1.TypeMeta `json:",inline"`
// 	// Standard object's metadata.
// 	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
// 	// +optional
// 	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
// 	Spec              NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
// 	Status            NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
// }
//
// type NodeSpec struct {
// 	PodCIDR       string
// 	Unschedulable bool
// 	// Simple custom slice if you want to support basic scheduling restrictions
// 	Taints []Taint
// }
//
// type Taint struct {
// 	// Required. The taint key to be applied to a node.
// 	Key string `json:"key" protobuf:"bytes,1,opt,name=key"`
// 	// The taint value corresponding to the taint key.
// 	// +optional
// 	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
// 	// Required. The effect of the taint on pods
// 	// that do not tolerate the taint.
// 	// Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
// 	Effect TaintEffect `json:"effect" protobuf:"bytes,3,opt,name=effect,casttype=TaintEffect"`
// 	// TimeAdded represents the time at which the taint was added.
// 	// +optional
// 	TimeAdded *metav1.Time `json:"timeAdded,omitempty" protobuf:"bytes,4,opt,name=timeAdded"`
// }
//
// type NodeCondition struct {
// 	// Type of node condition.
// 	Type NodeConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeConditionType"`
// 	// Status of the condition, one of True, False, Unknown.
// 	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
// 	// Last time we got an update on a given condition.
// 	// +optional
// 	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty" protobuf:"bytes,3,opt,name=lastHeartbeatTime"`
// 	// Last time the condition transit from one status to another.
// 	// +optional
// 	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
// 	// (brief) reason for the condition's last transition.
// 	// +optional
// 	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
// 	// Human readable message indicating details about last transition.
// 	// +optional
// 	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
// }
//
// type NodeStatus struct {
// 	Capacity    map[string]string // e.g. {"cpu": "4", "memory": "16Gi"}
// 	Allocatable map[string]string
// 	Phase       string          // "Ready", "NotReady"
// 	Conditions  []NodeCondition // Simple type with Type and Status strings
// 	Addresses   []string        // e.g. ["192.168.1.50"]
// }
