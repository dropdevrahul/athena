package zeus

import (
	"encoding/json"
	"net/http"
)

const (
	JsonContentTypeHeader string = "Content-Type"
	JsonContentType       string = "application/json; charset=utf-8"
)

func Json(w http.ResponseWriter, p interface{}, status int,
	headers map[string]string) error {
	w.Header().Set(JsonContentTypeHeader, JsonContentType)

	for h, val := range headers {
		w.Header().Set(h, val)
	}

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	w.Write(b)

	return nil
}
