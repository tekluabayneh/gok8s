package noderesourece

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"slices"

	Types "github.com/tekluabayneh/gok8s/types"
	v1 "k8s.io/api/core/v1"
)

// cfg := spew.ConfigState{Indent: "  ", DisableMethods: true}
//
//	fmt.Println(len(nodeinfo))
//	cfg.Dump("len", len(nodeinfo))
type ScoringStrategy string

const (
	LeastAllocated     ScoringStrategy = "LeastAllocated"
	MostAllocated      ScoringStrategy = "MostAllocated"
	BalancedAllocation ScoringStrategy = "BalancedAllocation"
)

type scorePlugin struct {
	name   string
	weight int64
	score  int64
}

type Scorer struct {
	strategy ScoringStrategy
	plugins  []scorePlugin
}

func NewScorer(strategy ScoringStrategy) *Scorer {
	return &Scorer{strategy: strategy}
}

func (sc *Scorer) Score(po *v1.Pod, nodeinfo []*Types.NodeInfo) {
	scoredNodes := sc.scoreNodes(po, nodeinfo)
	chosenNode, err := sc.selectHost(scoredNodes)
	if err != nil {
		panic(err)
	}

	fmt.Println("the best node", chosenNode)
}

func (sc *Scorer) scoreLeastAllocated(pod *v1.Pod, node *Types.NodeInfo) int64 {
	var usedmem int64
	var usedcpu int64

	for _, po := range node.Pods {
		for _, c := range po.Spec.Containers {
			usedcpu += c.Resources.Requests.Cpu().MilliValue()
			usedmem += c.Resources.Requests.Memory().Value()
		}
	}

	var reqpodmem int64
	var reqpodcpu int64
	for _, c := range pod.Spec.Containers {
		reqpodcpu += c.Resources.Requests.Cpu().MilliValue()
		reqpodmem += c.Resources.Requests.Memory().Value()
	}

	cpuFraction := (usedcpu + reqpodcpu) * 100 / node.Node.Status.Allocatable.Cpu().MilliValue()
	memFraction := (usedmem + reqpodmem) * 100 / node.Node.Status.Allocatable.Memory().Value()

	return 100 - (cpuFraction+memFraction)/2
}

func (sc *Scorer) scoreMostAllocated(pod *v1.Pod, node *Types.NodeInfo) int64 {
	var usedmem int64
	var usedcpu int64

	for _, po := range node.Pods {
		for _, c := range po.Spec.Containers {
			usedcpu += c.Resources.Requests.Cpu().MilliValue()
			usedmem += c.Resources.Requests.Memory().Value()
		}
	}

	var reqpodmem int64
	var reqpodcpu int64
	for _, c := range pod.Spec.Containers {
		reqpodcpu += c.Resources.Requests.Cpu().MilliValue()
		reqpodmem += c.Resources.Requests.Memory().Value()
	}

	cpuFraction := (usedcpu + reqpodcpu) * 100 / node.Node.Status.Allocatable.Cpu().MilliValue()
	memFraction := (usedmem + reqpodmem) * 100 / node.Node.Status.Allocatable.Memory().Value()

	return (cpuFraction + memFraction) / 2
}

func (sc *Scorer) scoreBalancedAllocation(pod *v1.Pod, node *Types.NodeInfo) int64 {
	var reqpodcpu int64
	var reqpodmem int64

	for _, c := range pod.Spec.Containers {
		reqpodcpu += c.Resources.Requests.Cpu().MilliValue()
		reqpodmem += c.Resources.Requests.Memory().Value()
	}

	var usedpodcpu int64
	var usedpodmem int64

	for _, po := range node.Pods {
		for _, c := range po.Spec.Containers {
			usedpodcpu += c.Resources.Requests.Cpu().MilliValue()
			usedpodmem += c.Resources.Requests.Memory().Value()
		}
	}
	cpuFraction := (usedpodcpu + reqpodcpu) * 100 / 2 / node.Node.Status.Allocatable.Cpu().MilliValue()
	memFraction := (usedpodmem + reqpodmem) * 100 / 2 / node.Node.Status.Allocatable.Memory().Value()

	diff := cpuFraction - memFraction
	if diff < 0 {
		diff = -diff // slip to positive if value end up negative
	}
	return 100 - diff
}

func (sc *Scorer) scoreResourceAllocation(pod *v1.Pod, node *Types.NodeInfo) int64 {
	switch sc.strategy {
	case MostAllocated:
		return sc.scoreMostAllocated(pod, node)
	case LeastAllocated:
		return sc.scoreLeastAllocated(pod, node)
	case BalancedAllocation:
		return sc.scoreBalancedAllocation(pod, node)
	default:
		return sc.scoreLeastAllocated(pod, node)
	}
}

func (sc *Scorer) scoreNodeAffinity(pod *v1.Pod, node *Types.NodeInfo) int64 {
	if pod.Spec.Affinity == nil || pod.Spec.Affinity.NodeAffinity == nil {
		return 0
	}
	terms := pod.Spec.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution
	var total int64
	for _, term := range terms {
		matched := true
		for _, expr := range term.Preference.MatchExpressions {
			val, ok := node.Node.Labels[expr.Key]
			if !ok || !slices.Contains(expr.Values, val) {
				matched = false
				break
			}
		}
		if matched {
			total += int64(term.Weight)
		}
	}
	return total
}

func (sc *Scorer) scoreNodes(pod *v1.Pod, feasible []*Types.NodeInfo) map[string]int64 {
	score := make(map[string]int64)
	for _, node := range feasible {
		scoreVal := sc.scoreResourceAllocation(pod, node)
		weightVal := sc.scoreNodeAffinity(pod, node)
		total := int(weightVal) + int(scoreVal)
		score[node.Node.Name] = int64(total)

	}
	return score
}

func (sc *Scorer) selectHost(scores map[string]int64) (string, error) {
	if len(scores) == 0 {
		return "", errors.New("no feasible nodes")
	}

	best := int64(math.MinInt)
	var tied []string

	for name, score := range scores {
		if score > best {
			best = score
			tied = []string{name}
		} else if score == best {
			tied = append(tied, name)
		}
	}

	return tied[rand.Intn(len(tied))], nil
}
