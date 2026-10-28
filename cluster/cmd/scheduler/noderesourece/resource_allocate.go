package noderesourece

import (
	"context"

	Types "github.com/tekluabayneh/gok8s/types"
	v1 "k8s.io/api/core/v1"
)

// Pressure conditions. MemoryPressure, DiskPressure and PIDPressure show up in node conditions, and real clusters turn them into taints.
// ============================================================
// 3. SCORE (only runs on nodes that passed Filter, ranks them)
// ============================================================

// Node affinity (preferred rules): preferredDuringSchedulingIgnoredDuringExecution.
// Resource allocation strategy: least allocated / most allocated / balanced.
//   This is the file you pasted.

type resourceAllocationScorer struct {
	Name                                          string
	enableInPlacePodVerticalScaling               bool
	enablePodLevelResources                       bool
	enableDRAExtendedResource                     bool
	enableInPlacePodLevelResourcesVerticalScaling bool
	// used to decide whether to use Requested or NonZeroRequested for
	// cpu and memory.
	useRequested bool
	scorer       func(requested, allocated, allocatable []int64) int64
	// resources    []config.ResourceSpec
}

// pattern must be used
// strategy parretn
// factory pattern
func (r *resourceAllocationScorer) score(
	ctx context.Context,
	pod *v1.Pod,
	nodeInfo Types.NodeInfo,
	podRequests []int64,
	// draPreScoreState *draPreScoreState,
) (int64, *bool) {
	var ff bool = false
	return 23, &ff
}

func (r *resourceAllocationScorer) calculateResourceAllocatableRequest() {
}

func (r *resourceAllocationScorer) calculatePodResourceRequest() {
}

func (r *resourceAllocationScorer) calculateNodeAllocatableRequest() {
}

func (r *resourceAllocationScorer) calculatePodResourceRequestList() {
}
