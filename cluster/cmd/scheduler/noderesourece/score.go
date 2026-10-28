package noderesourece

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	Types "github.com/tekluabayneh/gok8s/types"
)

func Score(nodeinfo []*Types.NodeInfo) {
	cfg := spew.ConfigState{Indent: "  ", DisableMethods: true}
	fmt.Println(len(nodeinfo))
	cfg.Dump("len", len(nodeinfo))
}
