package decoder

import (
	"encoding/json"
	"net/http"

	"github.com/tekluabayneh/gok8s/utils"
)

func Encoder(w http.ResponseWriter, status int, Msg any) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(Msg); err != nil {
		utils.Log().WithGroup("Encoder").Debug("Encoder failed to encode", "err", err)
		panic(err)
	}
}
