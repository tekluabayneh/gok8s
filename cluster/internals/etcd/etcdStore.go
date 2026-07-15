package etcd

import (
	"context"
	"time"

	"github.com/tekluabayneh/gok8s/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd() (*clientv3.Client, error) {
	Client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err == context.DeadlineExceeded {
		utils.Log().WithGroup("InitEtcd").Info("context timeout")
		return nil, err
	}

	if err != nil {
		utils.Log().WithGroup("InitEtcd").Debug("context timeout", "err", err)
		return nil, err
	}

	return Client, nil
}
