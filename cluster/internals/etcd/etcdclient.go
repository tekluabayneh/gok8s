package etcd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tekluabayneh/gok8s/config"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetEtcd(ctx context.Context, client *clientv3.Client, conf config.Pod) (*clientv3.GetResponse, error) {
	prefix := BuildKey(conf.Kind, conf.Metadata.Namespace, conf.Metadata.Name)

	// TODO
	// before string yaml value change to string
	val, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	_, err = client.Put(ctx, prefix, string(val))
	if err != nil {
		return nil, handlerEtcdError(err)
	}

	res, err := client.Get(ctx, prefix)
	fmt.Println("this is the value that retrived")
	if err != nil {
		return nil, handlerEtcdError(err)
	}

	return res, nil
}

func StoreEtcd(ctx context.Context, client *clientv3.Client, name, namespace, value string, kind string) (*clientv3.PutResponse, error) {
	prefix := BuildKey(kind, namespace, name)
	res, err := client.Put(ctx, prefix, value)
	if err != nil {
		return nil, handlerEtcdError(err)
	}

	return res, nil
}

func UpdateEtcd(ctx context.Context, client *clientv3.Client, name string, value string) (*clientv3.PutResponse, error) {
	res, err := client.Put(ctx, name, value)
	if err != nil {
		return nil, handlerEtcdError(err)
	}

	return res, nil
}

func DeleteEtcd(ctx context.Context, client *clientv3.Client, name string) (*clientv3.DeleteResponse, error) {
	res, err := client.Delete(ctx, name)
	if err != nil {
		return nil, handlerEtcdError(err)
	}

	return res, nil
}

func handlerEtcdError(err error) error {
	if clientv3.IsConnCanceled(err) {
		fmt.Println()
	}

	if err != nil {
		if clientv3.IsConnCanceled(err) {
			fmt.Println("gRPC client connection is closed", err)
		} else if err == context.Canceled {
			fmt.Println("ctx is canceled by another routine", err)
		} else if err == context.DeadlineExceeded {
			fmt.Println("ctx is attached with a deadline and it exceeded", err)
		} else if err == rpctypes.ErrEmptyKey {
			fmt.Println("client-side error: key is not provided", err)
		} else if ev, ok := status.FromError(err); ok {
			code := ev.Code()
			if code == codes.DeadlineExceeded {
				fmt.Println("server-side context might have timed-out first (due to clock skew) while original client-side context is not timed-out yet", err)
			}
		} else {
			fmt.Println("bad cluster endpoints, which are not etcd servers", err)
		}
	}
	return err
}
