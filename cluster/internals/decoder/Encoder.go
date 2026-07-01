package decoder

import (
	"encoding/json"
	"net/http"
)

func Encoder(w http.ResponseWriter, status int, Msg any) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(Msg); err != nil {
		panic(err)
	}
}
