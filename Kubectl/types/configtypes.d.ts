export type KubectlconfigType = {
  apiVersion: string;
  kind: string;
  clusters: {
    name: string;
    cluster: {
      server: string;
      "certificate-authority-data"?: string;
      "insecure-skip-tls-verify"?: boolean;
    };
  }[];
  users: {
    name: string;
    user: {
      token?: string;
      "client-certificate-data"?: string;
      "client-key-data"?: string;
    };
  }[];
  contexts: {
    name: string;
    context: {
      cluster: string;
      user: string;
      namespace?: string;
    };
  }[];
  "current-context": string;
};



type ContainerStatus = {
  name: string
  state: {
    running?: {
      startedAt: string
    }
    waiting?: {
      reason?: string
      message?: string
    }
    terminated?: {
      exitCode: number
      reason?: string
      message?: string
      startedAt?: string
      finishedAt?: string
    }
  }
  ready: boolean
  restartCount: number
  image: string
  imageID: string
  containerID?: string
  started?: boolean
}
export type RsPodtype = {
  metadata: {
    name: string
    namespace: string
    uid: string
    resourceVersion: string
    generation: number,
    creationTimestamp: string
    labels: {},
    managedFields: []
  },
  spec: {
    volumes: [],
    containers: [],
    restartPolicy: string
    terminationGracePeriodSeconds: number,
    dnsPolicy: string
    serviceAccountName: string
    serviceAccount: string
    nodeName: string
    securityContext: {},
    schedulerName: string
    tolerations: [],
    priority: number,
    enableServiceLinks: boolean,
    preemptionPolicy: string
  },
  status: {
    observedGeneration: 1,
    phase: string
    conditions: [],
    hostIP: string
    hostIPs: [],
    podIP: string
    podIPs: [],
    startTime: string
    containerStatuses: ContainerStatus[],
    qosClass: string
  }
}

// export type { KubectlConfigType, RsPodtype }
