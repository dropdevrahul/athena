package athena

import (
	"encoding/json"
	"net/http"
)

const (
	// JsonContentTypeHeader http header for content type
	JsonContentTypeHeader string = "Content-Type"
	// JsonContentType http header value for json response/requests
	JsonContentType string = "application/json; charset=utf-8"
)

// Json returns a application/json response with the given payload p and http status code
// status
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
