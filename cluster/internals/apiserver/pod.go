package internals

import (
	"context"
	"fmt"

	"github.com/tekluabayneh/gok8s/config"
	"github.com/tekluabayneh/gok8s/internals/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdStore struct {
	Client *clientv3.Client
}

func (store *EtcdStore) GetPod(ctx context.Context, conf config.Pod) (string, error) {
	fmt.Println("this is the getPod object that handler logincs")
	// LIFECYCLE: the GetPod() handler itself will stay in the code segment till there is request comming
	// MEMORY: it won't go to the EITHER the Heap OR the Stack it state in the Code Segment
	// FLOW: it only run when the cpu get request and want to access this handler block of code form the code segment
	//
	// 1. Allocation: Code Segment and accessed by the cpu when instruction require it
	// 2. Concurrency: Built on a single thread; concurrently read by thousands of HTTP request goroutines safely.
	// 3. Layout: its just way of accessing EtcdStore database and only name, namespace, kind, context,
	// 4. Failure: If request fail it wont' panic or crash the server it only request server error message

	// res, err := etcd.GetEtcd(ctx, store.client, res)
	res, err := etcd.GetEtcd(ctx, store.Client, conf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	return "", nil
}

func (store *EtcdStore) CreatePod(ctx context.Context, pod string) error {
	// update yaml
	// create pods yaml
	// etc...
	return nil
}

func (store *EtcdStore) DeletePod(ctx context.Context, namespace, pod string) error {
	return nil
}
