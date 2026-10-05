export const fakePodsForTest = {
  apiVersion: 'v1',
  kind: 'PodList',
  metadata: {
    resourceVersion: '123456',
  },
  items: [
    {
      apiVersion: 'v1',
      kind: 'Pod',

      metadata: {
        name: 'nginx',
        namespace: 'default',
        uid: 'nginx-12345',
        resourceVersion: '123456',
        generation: 1,
        creationTimestamp: '2026-10-05T10:00:00Z',
        labels: {
          app: 'nginx',
        },
        managedFields: [],
      },

      spec: {
        volumes: [],
        containers: [
          {
            name: 'nginx',
            image: 'nginx:latest',
          },
        ],
        restartPolicy: 'Always',
        terminationGracePeriodSeconds: 30,
        dnsPolicy: 'ClusterFirst',
        serviceAccountName: 'default',
        serviceAccount: 'default',
        nodeName: 'kind-worker',
        securityContext: {},
        schedulerName: 'default-scheduler',
        tolerations: [],
        priority: 0,
        enableServiceLinks: true,
        preemptionPolicy: 'PreemptLowerPriority',
      },

      status: {
        observedGeneration: 1,
        phase: 'Running',

        conditions: [],

        hostIP: '172.18.0.2',
        hostIPs: [
          {
            ip: '172.18.0.2',
          },
        ],

        podIP: '10.244.0.5',
        podIPs: [
          {
            ip: '10.244.0.5',
          },
        ],

        startTime: '2026-10-05T10:00:05Z',

        containerStatuses: [
          {
            name: 'nginx',
            state: {
              running: {
                startedAt: '2026-10-05T10:00:10Z',
              },
            },
            ready: true,
            restartCount: 0,
            image: 'nginx:latest',
            imageID: 'docker-pullable://nginx@sha256:123456',
            containerID: 'containerd://abc123',
            started: true,
          },
        ],

        qosClass: 'BestEffort',
      },
    },

    {
      apiVersion: 'v1',
      kind: 'Pod',

      metadata: {
        name: 'redis',
        namespace: 'default',
        uid: 'redis-67890',
        resourceVersion: '123457',
        generation: 1,
        creationTimestamp: '2026-10-05T10:05:00Z',
        labels: {
          app: 'redis',
        },
        managedFields: [],
      },

      spec: {
        volumes: [],
        containers: [
          {
            name: 'redis',
            image: 'redis:latest',
          },
        ],
        restartPolicy: 'Always',
        terminationGracePeriodSeconds: 30,
        dnsPolicy: 'ClusterFirst',
        serviceAccountName: 'default',
        serviceAccount: 'default',
        nodeName: 'kind-worker',
        securityContext: {},
        schedulerName: 'default-scheduler',
        tolerations: [],
        priority: 0,
        enableServiceLinks: true,
        preemptionPolicy: 'PreemptLowerPriority',
      },

      status: {
        observedGeneration: 1,
        phase: 'Pending',

        conditions: [],

        hostIP: '172.18.0.2',
        hostIPs: [
          {
            ip: '172.18.0.2',
          },
        ],

        podIP: '',
        podIPs: [],

        startTime: '2026-10-05T10:05:05Z',

        containerStatuses: [
          {
            name: 'redis',
            state: {
              waiting: {
                reason: 'ContainerCreating',
                message: 'Container is waiting to start',
              },
            },
            ready: false,
            restartCount: 0,
            image: 'redis:latest',
            imageID: '',
            started: false,
          },
        ],

        qosClass: 'BestEffort',
      },
    },
  ],
}
