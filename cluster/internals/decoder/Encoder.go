package decoder

import (
	"encoding/json"
	"io"
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

func DecodeList[T any](r io.Reader) (*T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}
