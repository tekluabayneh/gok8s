package etcd

import (
	"fmt"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

// TODO
// implement first in last out queue
// i need way to do like CRUD opration in queue
type IDeltaFIFO interface {
	Add(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64)
	Delete(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64)
	Update(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64)
	Sync(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64)
}

type DeltaType string

const (
	Added   DeltaType = "Added"
	Deleted DeltaType = "Update"
	Updated DeltaType = "Deleted"
	Sync    DeltaType = "Sync"
)

type Delta struct {
	Type   DeltaType
	Object interface{}
}

type DeltaFIFO struct {
	queue []string
	item  map[string]Delta
}

func (Del *DeltaFIFO) Add(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64) {
	fmt.Printf("Type: %s, Key: %s, Revision: %d\n", EventType, EventKey, EventModeRevision)

	DeltaFIFO{
		queue: []string{
			"EventKey",
		},

		item: map[string]string{
			"EventKey": {
				{
					Type:   Added,
					Object: "EventKey",
				},
			},
		},
	}
}

func (Del *DeltaFIFO) Delete(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64) {
	fmt.Printf("Type: %s, Key: %s, Revision: %d\n", EventType, EventKey, EventModeRevision)
}

func (Del *DeltaFIFO) Update(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64) {
	fmt.Printf("Type: %s, Key: %s, Revision: %d\n", EventType, EventKey, EventModeRevision)
}

func (Del *DeltaFIFO) Sync(EventType mvccpb.Event_EventType, EventKey []byte, EventModeRevision int64) {
	fmt.Printf("Type: %s, Key: %s, Revision: %d\n", EventType, EventKey, EventModeRevision)
}
