package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/tekluabayneh/gok8s/config"
)

func RegNode() {
	kl := Kub{}
	newk, err := kl.RegisterNode()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	json, err := json.Marshal(newk)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(json)
	cli, req := config.CreatApi(ctx, "POST", "/api/v1/nodes", body)
	res, err := cli.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("res", res)
}
