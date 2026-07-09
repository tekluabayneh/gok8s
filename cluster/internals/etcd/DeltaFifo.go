package etcd

import (
	"fmt"
	"sync"
)

// TODO
// implement first in last out queue
// i need way to do like CRUD opration in queue
type DeltaType string

const (
	Added   DeltaType = "Added"
	Updated DeltaType = "Updated"
	Deleted DeltaType = "Deleted"
	Sync    DeltaType = "Sync"
)

type IDeltaFIFO interface {
	Add(EventKey, EvenKeyVal []byte, EventModeRevision int64) error
	Delete(EventKey, EvenKeyVal []byte, EventModeRevision int64) error
	Update(EventKey, EvenKeyVal []byte, EventModeRevision int64) error
	Sync(EventKey []byte, EventModeRevision int64)
	Pop(any) (interface{}, error)
	List() any
}

type Delta struct {
	Type   DeltaType
	Object string
}

type DeltaFIFO struct {
	Item        map[string][]Delta
	Queue       []string
	LookUpQueue map[string]bool
	lock        sync.RWMutex
	closed      bool
}

var Fifo IDeltaFIFO = &DeltaFIFO{
	Item:        make(map[string][]Delta),
	Queue:       []string{},
	LookUpQueue: make(map[string]bool),
	closed:      false,
}

func (d *DeltaFIFO) close() bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.closed = true
	return d.closed
}

func (d *DeltaFIFO) queueActionInternalLocked(EventKey, EvenKeyVal []byte, EvenType DeltaType) error {
	key := string(EventKey)

	if key == "" {
		return fmt.Errorf("event key missing %s key", key)
	}

	if EvenType == Deleted {
		if ok := d.LookUpQueue[key]; !ok {
			return fmt.Errorf("couldn't delete, item does not exist with key of %s", key)
		}
	}

	delta := Delta{
		Type:   EvenType,
		Object: string(EvenKeyVal),
	}

	if ok := d.LookUpQueue[key]; !ok {
		d.Queue = append(d.Queue, key)
		d.LookUpQueue[key] = true
	}
	d.Item[key] = append(d.Item[key], delta)
	fmt.Println("Queue", d.Queue)
	fmt.Println("Item", d.Item)

	return nil
}

func (d *DeltaFIFO) Add(EventKey, EvenKeyVal []byte, EventModeRevision int64) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.queueActionInternalLocked(EventKey, EvenKeyVal, Added)
}

func (d *DeltaFIFO) Update(EventKey, EvenKeyVal []byte, EventModeRevision int64) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.queueActionInternalLocked(EventKey, EvenKeyVal, Updated)
}

func (d *DeltaFIFO) Delete(EventKey, EvenKeyVal []byte, EventModeRevision int64) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.queueActionInternalLocked(EventKey, EvenKeyVal, Deleted)
}

func (d *DeltaFIFO) Sync(EventKey []byte, EventModeRevision int64) {
	fmt.Printf("sync... called")
	fmt.Println("Item", d.Item)
	fmt.Println("Queue", d.Queue)
}

func (d *DeltaFIFO) Pop(procces any) (interface{}, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	for {
		if len(d.Queue) == 0 {
			if d.close() {
				fmt.Println("Queue is empty closed ")
			}
			return nil, fmt.Errorf("queue is empty")
		}
	}
}

func (d *DeltaFIFO) List() any {
	for _, Qu := range d.Queue {
		fmt.Printf("list of queue %s\n", Qu)
	}

	for _, Qu := range d.Item {
		for _, Del := range Qu {
			fmt.Printf("Delta Type %s and Delta object %s\n", Del.Type, Del.Object)
		}
	}
	return nil
}

///
// once the kubelet or schduler open long connectin to the keube Api  ther is channle based on resource type like pod node and then since the connection won't terminate it will wait till somethin get drop to that channel and when ever update add or delte added to that cahannle it stream it to them to those who needs that or leastning that moment
//
//
//so watch Event -> Added -> Queeu -> Pop -> EventLoop -> proccess -> resouce type chanlle like PodCahnnel -> Send to those who are lesning actively like if channel is emtpy just con't close the conneciton just wait ther ther is somethin gto send form the channel but not only one send and fisnhe prodcast it to all those lestning for that event
//
//
////etcd watch event (Added)
//   → Reflector.Add()
//   → DeltaFIFO.Queue
//   → processLoop calls Pop() (blocks until something's there)
//   → process(item) runs           ← this IS your "EventLoop -> process" step
//   → process() does the broadcast: sends the delta to the right resource-type
//     channel (your "PodChannel" idea)
//   → WatchBroadcaster fans that single delta out to every subscriber
//     channel currently registered for that resource type
//   → each watcher's HTTP handler (blocked, waiting) wakes up and writes
//     it out over its own connection
////
