package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type PodStore interface {
	GetPod(ctx context.Context, namespace, name string) (string, error)
	CreatePod(ctx context.Context, pod string) error
	DeletePod(ctx context.Context, namespace string, pod string) error
}

type PodHnalder struct {
	Store PodStore
}

func (p *PodHnalder) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if err := p.Store.CreatePod(ctx, "name"); err != nil {
		fmt.Println(err)
	}
}

func (p *PodHnalder) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is get pdo handler")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if _, err := p.Store.GetPod(ctx, "name", "namespace"); err != nil {
		fmt.Println(err)
	}
}

func (p *PodHnalder) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if err := p.Store.DeletePod(ctx, "namespace", "pod"); err != nil {
		fmt.Println(err)
	}
}
