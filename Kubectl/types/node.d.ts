
export interface ObjectMeta {
  name: string;
  uid: string;
  resourceVersion: string;
  creationTimestamp: string;
  labels?: Record<string, string>;
  annotations?: Record<string, string>;
  managedFields?: ManagedFieldsEntry[];
}

export interface ManagedFieldsEntry {
  manager: string;
  operation: string;
  apiVersion: string;
  time: string;
  fieldsType: string;
  fieldsV1: Record<string, unknown>;
  subresource?: string;
}

export interface NodeTaint {
  key: string;
  value?: string;
  effect: string;
  timeAdded?: string;
}

export interface NodeSpec {
  podCIDR?: string;
  podCIDRs?: string[];
  providerID?: string;
  taints?: NodeTaint[];
  unschedulable?: boolean;
}

export interface NodeCondition {
  type: string;
  status: string;
  lastHeartbeatTime: string;
  lastTransitionTime: string;
  reason: string;
  message: string;
}

export interface NodeAddress {
  type: string;
  address: string;
}

export interface NodeDaemonEndpoints {
  kubeletEndpoint: {
    Port: number;
  };
}

export interface NodeSystemInfo {
  machineID: string;
  systemUUID: string;
  bootID: string;
  kernelVersion: string;
  osImage: string;
  containerRuntimeVersion: string;
  kubeletVersion: string;
  kubeProxyVersion: string;
  operatingSystem: string;
  architecture: string;
  swap?: {
    capacity: number;
  };
}

export interface ContainerImage {
  names: string[];
  sizeBytes: number;
}

export interface NodeRuntimeHandler {
  name: string;
  features: {
    recursiveReadOnlyMounts?: boolean;
    userNamespaces?: boolean;
  };
}

export interface NodeFeatures {
  supplementalGroupsPolicy?: boolean;
}

// Resource quantities (cpu, memory, pods, etc.) are strings in the k8s API,
// e.g. "12", "7786724Ki" — parse them yourself where you need numbers.
export type ResourceList = Record<string, string>;

export interface NodeStatus {
  capacity: ResourceList;
  allocatable: ResourceList;
  conditions: NodeCondition[];
  addresses: NodeAddress[];
  daemonEndpoints: NodeDaemonEndpoints;
  nodeInfo: NodeSystemInfo;
  images: ContainerImage[];
  runtimeHandlers?: NodeRuntimeHandler[];
  features?: NodeFeatures;
}

export interface Node {
  kind: 'Node';
  apiVersion: string;
  metadata: ObjectMeta;
  spec: NodeSpec;
  status: NodeStatus;
}
