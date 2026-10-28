package handlers

import (
	"context"

	"github.com/tekluabayneh/gok8s/config"
	"github.com/tekluabayneh/gok8s/internals/decoder"
	corev1 "k8s.io/api/core/v1"
)

func HandlerNodeStatus() []corev1.Node {
	ctx := context.Background()
	cli, req := config.CreatApi(ctx, "GET", "/api/v1/nodes", nil)
	res, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	nodeList, err := decoder.DecodeList[corev1.NodeList](res.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(nodeList)
	// TODO:
	// after body read extract the node status needs
	return nodeList.Items
}
