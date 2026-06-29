package etcd

import "fmt"

type ResourceType string

const (
	// Workload Layer
	ResourcePod         ResourceType = "pods"
	ResourceDeployment  ResourceType = "deployments"
	ResourceReplicaSet  ResourceType = "replicasets"
	ResourceStatefulSet ResourceType = "statefulsets"
	ResourceDaemonSet   ResourceType = "daemonsets"
	ResourceJob         ResourceType = "jobs"
	ResourceCronJob     ResourceType = "cronjobs"

	// Networking Layer
	ResourceService       ResourceType = "services"
	ResourceIngress       ResourceType = "ingresses"
	ResourceEndpoint      ResourceType = "endpoints"
	ResourceNetworkPolicy ResourceType = "networkpolicies"

	// Configuration & Storage Layer
	ResourceConfigMap             ResourceType = "configmaps"
	ResourceSecret                ResourceType = "secrets"
	ResourcePersistentVolumeClaim ResourceType = "persistentvolumeclaims"
	ResourcePersistentVolume      ResourceType = "persistentvolumes"
	ResourceStorageClass          ResourceType = "storageclasses"

	// Security & RBAC Layer
	ResourceServiceAccount     ResourceType = "serviceaccounts"
	ResourceRole               ResourceType = "roles"
	ResourceRoleBinding        ResourceType = "rolebindings"
	ResourceClusterRole        ResourceType = "clusterroles"
	ResourceClusterRoleBinding ResourceType = "clusterrolebindings"

	// System & Extension Layer
	ResourceNode                     ResourceType = "nodes"
	ResourceEvent                    ResourceType = "events"
	ResourceCustomResourceDefinition ResourceType = "customresourcedefinitions"
)

// [Root] ──► [Resource Type] ──► [Namespace Name] ──► [Resource Nam]
// "registry" "pods or configmaps" "colab-app" "colab-app-pods"

func BuildKey(kind, namespace, name string) string {
	if namespace == "" {
		return fmt.Sprintf("gok8s/%s/%s/", kind, name)
	}
	return fmt.Sprintf("gok8s/%s/%s/%s/", kind, namespace, name)
}

// used for universal staff like nodes and other cluster wide staff
func BuildPrefix(kind, namespace string) string {
	if namespace == "" {
		return fmt.Sprintf("gok8s/%s/", kind)
	}
	return fmt.Sprintf("gok8s/%s/%s/", kind, namespace)
}
