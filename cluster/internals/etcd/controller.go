package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Queue interface{}

type (
	PopProcessFunc func(obj interface{}, isInInitialList bool) error
	ProcessFunc    func(obj interface{}, isInInitialList bool) error
)

type Config struct {
	Queue        interface{}
	ListWatcher  int
	Process      ProcessFunc
	ProcessBatch func()
}

type controller struct {
	config           Config
	reflector        any // need to be change to real reflector
	reflectorWithMux sync.RWMutex
	clock            any // need to be changed to real time.clock func
}

type Controller interface {
	RunWithContext(ctx context.Context)
	HasSynched()
	Run(chStop <-chan int) error
	HassynchedCher()
	LastSyncedResouceVersion() string
}

func New(c *Config) *controller {
	ctrl := &controller{
		config: *c,
		clock:  time.Second * 4, // need to be change to real clock function
	}

	return ctrl
}

func (c *Config) Run() error {
	return fmt.Errorf("fake error")
}

func (c *Config) RunWithContext(ctx context.Context) {
}

func (c *Config) HasSynched() {
}

func (c *Config) HassynchedCher() {
}

func (c *Config) LastSyncedResouceVersion() string {
	return ""
}

func ProcessLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():

		default:
			fmt.Errorf("default error")
		}
	}
}
