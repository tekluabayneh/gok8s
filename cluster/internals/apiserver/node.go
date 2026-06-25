package internals

import "context"

func (p *EtcdStore) GetNode(ctx context.Context, name string) {
	p.client.Get(ctx, name)
}

func (p *EtcdStore) RegisterNode(ctx context.Context, name string, value string) {
	p.client.Put(ctx, name, value)
}
