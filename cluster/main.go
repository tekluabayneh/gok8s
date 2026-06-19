package main

import apiserver "github.com/tekluabayneh/gok8s/cmd/apiServer"

func main() {
	apiserver.AppAPIServerNew().APIServerStart()
}
