package etcd

import (
	"fmt"
)

// TODO
// implement first in last out queue
// i need way to do like CRUD opration in queue
type DeltaType string

const (
	Added   DeltaType = "Added"
	Deleted DeltaType = "Update"
	Updated DeltaType = "Deleted"
	Sync    DeltaType = "Sync"
)

type IDeltaFIFO interface {
	Add(EventKey []byte, EventModeRevision int64)
	Delete(EventKey []byte, EventModeRevision int64)
	Update(EventKey []byte, EventModeRevision int64)
	Sync(EventKey []byte, EventModeRevision int64)
	Pop()
	List() []interface{}
}

type Delta struct {
	Type   DeltaType
	Object interface{}
}

type DeltaFIFO struct {
	Item  map[string][]Delta
	Queue []string
}

var Fifo IDeltaFIFO = &DeltaFIFO{
	Item:  make(map[string][]Delta),
	Queue: []string{},
}

func (d *DeltaFIFO) Add(EventKey []byte, EventModeRevision int64) {
	fmt.Printf(" Key: %s, Revision: %d\n", EventKey, EventModeRevision)
	key := string(EventKey)
	d.Queue = append(d.Queue, key)
	delta := Delta{
		Type:   "Added",
		Object: EventKey,
	}
	d.Item[key] = append(d.Item[key], delta)

	fmt.Printf("Add called\n")

	fmt.Println(d.Item)
	fmt.Println(d.Queue)
}

func (Del *DeltaFIFO) Delete(EventKey []byte, EventModeRevision int64) {
	fmt.Printf("Key: %s, Revision: %d\n", EventKey, EventModeRevision)
	fmt.Printf("delete called")
}

func (d *DeltaFIFO) Update(EventKey []byte, EventModeRevision int64) {
	fmt.Println(d.Item)
	fmt.Println(d.Queue)

	fmt.Printf("update called")
}

func (d *DeltaFIFO) Sync(EventKey []byte, EventModeRevision int64) {
	fmt.Printf("sync... called")
}

func (d *DeltaFIFO) Pop() {
	fmt.Printf("Pop... called")
}

func (d *DeltaFIFO) List() []interface{} {
	fmt.Printf("list of Delta's")
	return nil
}
