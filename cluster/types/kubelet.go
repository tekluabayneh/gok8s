package types

import (
	"net/http"
	"time"

	v1 "k8s.io/api/core/v1"
)

type NodeName string

type Taint struct {
	Key    string
	Value  string
	Effect string
}

type ObjectReference struct {
	Kind      string
	Namespace string
	Name      string
}

// Configuration

type KubeletConfiguration struct {
	RootDirectory string
	MaxPods       int
}

// Component contracts.

type (
	RuntimeCache               interface{}
	Client                     interface{}
	PodWorkers                 interface{}
	PodManager                 interface{}
	EvictionManager            interface{}
	CommandRunner              interface{}
	Cadvisor                   interface{}
	ServiceLister              interface{}
	NodeLister                 interface{}
	VolumePluginManager        interface{}
	ProbeManager               interface{}
	ProbeResultsManager        interface{}
	EventRecorder              interface{}
	StatusManager              interface{}
	VolumeManager              interface{}
	ContainerRuntime           interface{}
	StreamingRuntime           interface{}
	RuntimeService             interface{}
	LeaseController            interface{}
	PodLifecycleEventGenerator interface{}
	PodCache                   interface{}
	OOMWatcher                 interface{}
	MountInterface             interface{}
	ContainerManager           interface{}
)

// Kubelet manages node-level container and Pod operations.
type Kubelet struct {
	// Kubelet configuration.
	KubeletConfiguration KubeletConfiguration

	// Host identity.
	Hostname           string
	HostnameOverridden bool
	NodeName           NodeName

	// Runtime cache.
	RuntimeCache RuntimeCache

	// Kubernetes API clients.
	KubeClient      Client
	HeartbeatClient Client

	// Filesystem.
	RootDirectory string

	// Pod processing.
	PodWorkers PodWorkers
	PodManager PodManager

	// Resource eviction.
	EvictionManager EvictionManager

	// HTTP log server.
	LogServer http.Handler

	// Container command execution.
	Runner CommandRunner

	// Node monitoring.
	Cadvisor Cadvisor

	// Node registration.
	RegisterNode       func() (v1.Node, error)
	RegisterWithTaints []Taint

	// Informer listers and synchronization.
	ServiceLister    ServiceLister
	ServiceHasSynced func() bool
	NodeLister       NodeLister
	NodeHasSynced    func() bool

	// Node labels.
	NodeLabels map[string]string

	// Volume plugins.
	VolumePluginMgr VolumePluginManager

	// Health probes.
	ProbeManager     ProbeManager
	LivenessManager  ProbeResultsManager
	ReadinessManager ProbeResultsManager
	StartupManager   ProbeResultsManager

	// Events.
	Recorder EventRecorder

	// Managers.
	StatusManager StatusManager
	VolumeManager VolumeManager

	// Reference to this node.
	NodeRef *ObjectReference

	// Container runtime.
	ContainerRuntime ContainerRuntime
	StreamingRuntime StreamingRuntime
	RuntimeService   RuntimeService

	// Node status reporting.
	NodeStatusUpdateFrequency time.Duration
	NodeStatusReportFrequency time.Duration
	LastStatusReportTime      time.Time

	// Node lease controller.
	NodeLeaseController LeaseController

	// Pod lifecycle event generators.
	Pleg        PodLifecycleEventGenerator
	EventedPleg PodLifecycleEventGenerator

	// Cached Pod statuses.
	PodCache PodCache

	// Out-of-memory event watcher.
	OOMWatcher OOMWatcher

	// Volume mounting.
	Mounter MountInterface

	// Non-runtime container management.
	ContainerManager ContainerManager

	// Maximum Pods allowed on this node.
	MaxPods int
}

// type Kubelet struct {
// 	// kubelet config 对象
// 	kubeletConfiguration kubeletconfiginternal.KubeletConfiguration
//
// 	// hostname is the hostname the kubelet detected or was given via flag/config
// 	hostname string
// 	// hostnameOverridden indicates the hostname was overridden via flag/config
// 	hostnameOverridden bool
// 	// 节点名称
// 	nodeName types.NodeName
// 	// 运行时缓存，用于管理容器运行时的状态与信息
// 	runtimeCache kubecontainer.RuntimeCache
// 	// 客户端
// 	kubeClient      clientset.Interface
// 	heartbeatClient clientset.Interface
// 	// kubelet根目录
// 	rootDirectory string
// 	// podWorkers 用于处理 Pod 事件并同步 Pod 的工作器，重要对象
// 	// 处理每一个 Pod 的更新事件，创建 删除 更新等
// 	podWorkers PodWorkers
//
// 	// 提供接口来存储和访问 Pod 信息
// 	podManager kubepod.Manager
//
// 	// 当 Node 资源不足时，达到 evict 的策略时，会把 Pod 从 Node 上驱逐，来保证 Node 的稳定性
// 	evictionManager eviction.Manager
//
// 	// Optional, defaults to /logs/ from /var/log
// 	logServer http.Handler
// 	// 用于执行容器命令的接口
// 	runner kubecontainer.CommandRunner
//
// 	// cAdvisor 集成在 kubelet 中，收集节点和容器的监控信息
// 	cadvisor cadvisor.Interface
// 	// 是否注册节点
// 	registerNode bool
// 	// 注册时加入的污点列表
// 	registerWithTaints []v1.Taint
//
// 	// informer Lister 相关
// 	serviceLister    serviceLister
// 	serviceHasSynced cache.InformerSynced
// 	nodeLister       corelisters.NodeLister
// 	nodeHasSynced    cache.InformerSynced
// 	// a list of node labels to register
// 	nodeLabels map[string]string
// 	// Volume plugins.
// 	volumePluginMgr *volume.VolumePluginMgr
// 	// 健康检查或探针管理器
// 	probeManager     prober.Manager
// 	livenessManager  proberesults.Manager
// 	readinessManager proberesults.Manager
// 	startupManager   proberesults.Manager
//
// 	// 事件发送器
// 	recorder record.EventRecorder
//
// 	// 各种 Manager 管理器
//
// 	// statusManager
// 	statusManager status.Manager
// 	// volumeManager
// 	volumeManager volumemanager.VolumeManager
//
// 	// Reference to this node.
// 	nodeRef *v1.ObjectReference
//
// 	// Container runtime.
// 	containerRuntime kubecontainer.Runtime
// 	// Streaming runtime handles container streaming.
// 	streamingRuntime kubecontainer.StreamingRuntime
// 	// Container runtime service (needed by container runtime Start()).
// 	runtimeService internalapi.RuntimeService
//
// 	// 各种时间，主要用于上报操作
// 	nodeStatusUpdateFrequency time.Duration
// 	nodeStatusReportFrequency time.Duration
// 	lastStatusReportTime      time.Time
//
// 	// nodeLeaseController 租约控制器
// 	nodeLeaseController lease.Controller
//
// 	// pleg Manager
// 	pleg pleg.PodLifecycleEventGenerator
//
// 	// Evented PLEG Manager
// 	eventedPleg pleg.PodLifecycleEventGenerator
//
// 	// Store kubecontainer.PodStatus for all pods.
// 	podCache kubecontainer.Cache
//
// 	// Watcher of out of memory events.
// 	oomWatcher oomwatcher.Watcher
//
// 	// Mounter to use for volumes.
// 	mounter mount.Interface
// 	// Manager of non-Runtime containers.
// 	containerManager cm.ContainerManager
// 	// Maximum Number of Pods which can be run by this Kubelet
// 	maxPods int
// }
