package scheduler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	Types "github.com/tekluabayneh/gok8s/types"
)

const (
	Added    Types.EventType = "ADDED"
	Modified Types.EventType = "MODIFIED"
	Deleted  Types.EventType = "DELETED"
)

type Watcher struct {
	url    string
	client *http.Client
}

func NewWatcher(url string) *Watcher {
	return &Watcher{
		url:    url,
		client: &http.Client{},
	}
}

func (w *Watcher) Watch(ctx context.Context, events chan<- Types.WatchEvent) error {
	client := &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
	req, err := http.NewRequestWithContext(ctx, "GET", w.url+"?watch=true", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	for {
		var ev Types.WatchEvent
		if err := dec.Decode(&ev); err != nil {
			panic(err)
		}
		select {
		case events <- ev:
		case <-ctx.Done():
			// since w.Watch()  so when this case is called the function return and close() is called closing the channel
			// close(event) // even is channel
			return ctx.Err()
		}

	}
}
