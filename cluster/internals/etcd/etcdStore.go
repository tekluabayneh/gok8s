package etcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd() (*clientv3.Client, error) {
	Client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err == context.DeadlineExceeded {
		fmt.Println("context timeout")
		return nil, err
	}

	if err != nil {
		fmt.Println("context timeout")
		return nil, err
	}

	return Client, nil
}
