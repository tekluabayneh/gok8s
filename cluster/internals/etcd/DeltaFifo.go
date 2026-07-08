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
	Add(EventKey, EvenKeyVal []byte, EventModeRevision int64)
	Delete(EventKey []byte, EventModeRevision int64)
	Update(EventKey, EvenKeyVal []byte, EventModeRevision int64)
	Sync(EventKey []byte, EventModeRevision int64)
	Pop()
	List() []interface{}
}

type Delta struct {
	Type   DeltaType
	Object string
}

type DeltaFIFO struct {
	Item        map[string][]Delta
	Queue       []string
	LookUpQueue map[string]bool
}

var Fifo IDeltaFIFO = &DeltaFIFO{
	Item:        make(map[string][]Delta),
	Queue:       []string{},
	LookUpQueue: make(map[string]bool),
}

func (d *DeltaFIFO) Add(EventKey, EvenKeyVal []byte, EventModeRevision int64) {
	key := string(EventKey)

	delta := Delta{
		Type:   "Added",
		Object: string(EvenKeyVal),
	}

	if ok := d.LookUpQueue[key]; !ok {
		d.Queue = append(d.Queue, key)
		d.LookUpQueue[key] = true
	}
	d.Item[key] = append(d.Item[key], delta)
	fmt.Println("Queue", d.Queue)
	fmt.Println("Item", d.Item)
}

func (Del *DeltaFIFO) Delete(EventKey []byte, EventModeRevision int64) {
	fmt.Printf("Key: %s, Revision: %d\n", EventKey, EventModeRevision)
	fmt.Printf("delete called")
}

func (d *DeltaFIFO) Update(EventKey, EvenKeyVal []byte, EventModeRevision int64) {
	delta := Delta{
		Type:   "Updated",
		Object: string(EvenKeyVal),
	}

	d.Item[string(EventKey)] = append(d.Item[string(EventKey)], delta)

	fmt.Println("Item", d.Item)
	fmt.Println("Queue", d.Queue)
}

func (d *DeltaFIFO) Sync(EventKey []byte, EventModeRevision int64) {
	fmt.Printf("sync... called")
	fmt.Println("Item", d.Item)
	fmt.Println("Queue", d.Queue)
}

func (d *DeltaFIFO) Pop() {
	fmt.Printf("Pop... called")
	fmt.Println("Item", d.Item)
	fmt.Println("Queue", d.Queue)
}

func (d *DeltaFIFO) List() []interface{} {
	fmt.Printf("list of Delta's")
	fmt.Println("Item", d.Item)
	fmt.Println("Queue", d.Queue)

	return nil
}
