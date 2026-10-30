package scheduler

import (
	"context"
	"fmt"

	"github.com/tekluabayneh/gok8s/cmd/scheduler/handlers"
	"github.com/tekluabayneh/gok8s/cmd/scheduler/noderesourece"
	Types "github.com/tekluabayneh/gok8s/types"
	v1 "k8s.io/api/core/v1"
)

// TODO:
// assign pod to node
// send http reqeust to kube-api server
// check which node is best like which node score is higher
// and which node is typs is best for that prticular node
// READ the Scheduler demon documentation from k8s docs

// step
// must check node health
// much check node speci like cpu and memory capacity
type Scheduler struct {
	// node list, queue, etc.
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Run(ctx context.Context, events <-chan Types.WatchEvent) {
	for {
		select {
		case <-ctx.Done():
			return
		case ev, ok := <-events:
			if !ok {
				panic(ok)
			}

			Evaluator(ev.Object)
		}
	}
}

// call the filter and score funcito
func Evaluator(p v1.Pod) []*Types.NodeInfo {
	var SuitedNodeForThePod []*Types.NodeInfo
	nodelist := handlers.HandlerNodeStatus()
	scoreNode := noderesourece.NewScorer("least")
	f := noderesourece.NewFilterResource()

	// if p.Spec.NodeName != "" {
	// 	return []no // need to changed to the node it was assign
	// }

	for _, val := range nodelist {
		fesiableNode, reson := f.Filter(&Types.NodeInfo{Node: &val, Pods: []*v1.Pod{}}, p)
		SuitedNodeForThePod = append(SuitedNodeForThePod, fesiableNode)
		fmt.Println(reson)

	}
	scoreNode.Score(&p, SuitedNodeForThePod)
	return SuitedNodeForThePod
}
