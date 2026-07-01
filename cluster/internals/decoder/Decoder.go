package decoder

import (
	"encoding/json"
	"net/http"

	"github.com/tekluabayneh/gok8s/config"
)

func Decoder(r *http.Request) config.Pod {
	var jsonData config.Pod
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		panic(err)
	}
	return jsonData
}
