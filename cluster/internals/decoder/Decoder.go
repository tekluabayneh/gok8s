package decoder

import (
	"encoding/json"
	"net/http"

	"github.com/tekluabayneh/gok8s/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ptTypeObject[T any] interface {
	*T
	metav1.Object
}

func Decoder[T any, PT ptTypeObject[T]](r *http.Request) (PT, error) {
	jsonData := PT(new(T))
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		utils.Log().WithGroup("Debugger").Debug("Decoder failed to decode", "err", err)
		return nil, err
	}
	return jsonData, nil
}
