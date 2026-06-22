package internals

import (
	"context"
	"fmt"

	"github.com/tekluabayneh/gok8s/internals/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdStore struct {
	client *clientv3.Client
}

func (store *EtcdStore) GetPod(ctx context.Context, namespace, name string, kind etcd.ResourceType) (string, error) {
	fmt.Println("this is the getPod object that handler logincs")
	res, err := etcd.GetEtcd(ctx, store.client, name, namespace, kind)
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
