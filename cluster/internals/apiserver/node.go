package internals

import "context"

func (Store *EtcdStore) GetNode(ctx context.Context, name string) {
	Store.Client.Get(ctx, name)
}

func (Store *EtcdStore) RegisterNode(ctx context.Context, name string, value string) {
	Store.Client.Put(ctx, name, value)
}
