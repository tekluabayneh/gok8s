package decoder

import (
	"encoding/json"
	"net/http"

	"github.com/tekluabayneh/gok8s/config"
	"github.com/tekluabayneh/gok8s/utils"
)

func Decoder(r *http.Request) config.Pod {
	var jsonData config.Pod
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		utils.Log().WithGroup("Debugger").Debug("Decodeer failed to decode", "err", err)
		panic(err)
	}
	return jsonData
}
