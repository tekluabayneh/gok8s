package etcd

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO
//
//	use buildKey to build the key before passing path to the etcd and same as when performing any ectd staff
func GetEtcd(ctx context.Context, client *clientv3.Client, name, namespace string, kind ResourceType) (*clientv3.GetResponse, error) {
	prefix := BuildKey(kind, namespace, name)
	res, err := client.Get(ctx, prefix)
	if err != nil {
		return nil, handlerEtcdError(err)
	}
	return res, nil
}

func StoreEtcd(ctx context.Context, client *clientv3.Client, name, namespace, value string, kind ResourceType) (*clientv3.PutResponse, error) {
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
