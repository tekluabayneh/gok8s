package config

import (
	"time"
)

const (
	PodPending   PodPhase = "Pending"
	PodRunning   PodPhase = "Running"
	PodSucceeded PodPhase = "Succeeded"
	PodFailed    PodPhase = "Failed"
	PodUnknown   PodPhase = "Unknown"
)

type Pod struct {
	Kind       string     `json:"kind,inline"`
	APIVersion string     `json:"apiVersion"`
	Metadata   ObjectMeta `json:"metadata,omitempty"`
	Spec       PodSpec    `json:"spec,omitempty"`
	Status     PodStatus  `json:"status,omitempty"`
}

type PodStatus struct {
	Phase             PodPhase          `json:"phase,omitempty"`
	Conditions        []PodCondition    `json:"conditions,omitempty"`
	ContainerStatuses []ContainerStatus `json:"containerStatuses,omitempty"`
	PodIP             string            `json:"podIP,omitempty"`
	StartTime         *time.Time        `json:"startTime,omitempty"`
}

type (
	PodPhase     string
	PodCondition struct {
		Type   string `json:"type"`
		Status string `json:"status"`
	}
)

type ContainerStatus struct {
	Name    string `json:"name"`
	Ready   bool   `json:"ready"`
	Started *bool  `json:"started,omitempty"`
	Image   string `json:"image"`
}

type ObjectMeta struct {
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	UID               string            `json:"uid,omitempty"`
	CreationTimestamp time.Time         `json:"creationTimestamp,omitempty"`
	DeletionTimestamp *time.Time        `json:"deletionTimestamp,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
}

type TypeMeta struct {
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

type PodSpec struct {
	Containers      []Container         `json:"containers"`
	SecurityContext *PodSecurityContext `json:"securityContext,omitempty"`
}

type Container struct {
	Name            string                `json:"name"`
	Image           string                `json:"image"`
	Ports           []ContainerPort       `json:"ports,omitempty"`
	ReadinessProbe  *Probe                `json:"readinessProbe,omitempty"`
	Resources       *ResourceRequirements `json:"resources,omitempty"`
	SecurityContext *SecurityContext      `json:"securityContext,omitempty"`
}

type ContainerPort struct {
	ContainerPort int32  `json:"containerPort"`
	Protocol      string `json:"protocol,omitempty"`
}

type Probe struct {
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty"`
}

type HTTPGetAction struct {
	Path string `json:"path"`
	Port int32  `json:"port"`
}

type ResourceRequirements struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}

type SecurityContext struct {
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty"`
	Privileged               *bool `json:"privileged,omitempty"`
}

type PodSecurityContext struct {
	RunAsUser      *int64          `json:"runAsUser,omitempty"`
	RunAsGroup     *int64          `json:"runAsGroup,omitempty"`
	RunAsNonRoot   *bool           `json:"runAsNonRoot,omitempty"`
	SeccompProfile *SeccompProfile `json:"seccompProfile,omitempty"`
}

type SeccompProfile struct {
	Type string `json:"type"`
}
