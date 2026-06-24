package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tekluabayneh/gok8s/internals/etcd"
)

type PodStore interface {
	GetPod(ctx context.Context, namespace, name string, kind etcd.ResourceType) (string, error)
	CreatePod(ctx context.Context, pod string) error
	DeletePod(ctx context.Context, namespace string, pod string) error
}

type PodHnalder struct {
	Store PodStore
}

func (p *PodHnalder) Get(w http.ResponseWriter, r *http.Request) {
	// LIFECYCLE: the Get() handler itself will stay in the code segment till there is request comming
	// MEMORY: it won't go to the EITHER the Heap OR the Stack it state in the Code Segment
	// FLOW: it only run when the cpu get request and want to access this handler block of code form the code segment
	//
	// 1. Allocation: Code Segment and accessed by the cpu when instruction require it
	// 2. Concurrency: Built on a single thread; concurrently read by thousands of HTTP request goroutines safely.
	// 3. Layout: it's http.request and http.ResponseWriter layout accessing what is coming and respoing what will leave and accpept PodHnalder struct pointer
	// 4. Failure: If request fail it wont' panic or crash the server it only request server error message

	fmt.Println("this is get pdo handler")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	// LIFECYCLE: this context.WithTimeout get called when ever Get handler is called  and wait 5 second using go Runtime schduler and alert us when it finishes
	// MEMORY: it will scape to heap due to internal challen Done and cancel closure function
	// FLOW: it runs when the handler called and GO runtime will  and go tkae pointer stack to teh channle of Done so this get ckicked to heap
	//
	// 1. Allocation: moved to heap due to internal go runtime pointer and closure function
	// 2. Concurrency: multipe thread can access it and its' safely accessed by those goroutines and os Thread  and the internal channle will update safly to al cpu core if time finished
	// 3. Layout: it accept context and tie how long it will wait
	// 4. Failure: If request fail it wont' panic or crash the server it just propagates a cancellation signal message/error.

	defer cancel()

	if _, err := p.Store.GetPod(ctx, "name", "namespace", "configMaps"); err != nil {
		fmt.Println(err)
	}
}

func (p *PodHnalder) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if err := p.Store.CreatePod(ctx, "name"); err != nil {
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
