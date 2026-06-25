package handlers

import (
	"context"
	"net/http"
)

type NodeStore interface {
	GetNode(ctx context.Context, name string)
	RegisterNode(ctx context.Context, name string, value string)
}

type NodeHandler struct {
	Store NodeStore
}

// hard coded configuration
type Default struct {
	Controleplane string
}

//// TODO: POST /api/v1/nodes
//   - parse node info from request body
//   - validate (name can't be empty)
//   - write to etcd at key /registry/nodes/{name}
//   - return 201

// TODO: GET /api/v1/nodes
//   - read all keys under /registry/nodes/
//   - return as JSON array

func (p *NodeHandler) GetNodeHandler(w http.ResponseWriter, r *http.Request) {
	p.Store.GetNode(r.Context(), "fakename")
}

func (p *NodeHandler) RegisterNodeHandler(w http.ResponseWriter, r *http.Request) {
	p.Store.GetNode(r.Context(), "fakename")
}
