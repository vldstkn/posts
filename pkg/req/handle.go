package req

import (
	"net/http"
	"p1/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid[T](body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
	}
	return &body, nil
}
