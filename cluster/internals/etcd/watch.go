package etcd

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

go func watchPodEven(ctx context.Context, client *clientv3.Client) *clientv3.Event {
	// TODO
	// implement watch function taht watches for etcd db and kubectl and schduler will use this grpc based protocole to schdule and manage contianer
	// func WatchEtcd(ctx context.Context, client *clientv3.Client) {

	watcher := client.Watch(ctx, "gok8s/pods", clientv3.WithPrefix())
	for response := range watcher {
		for _, event := range response.Events {
			fmt.Printf("Type: %s, Key: %s, Revision: %d\n", event.Type, event.Kv.Key, event.Kv.ModRevision)
			return event
		}
	}
	return nil
}()


