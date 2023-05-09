package athena_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/dropdevrahul/athena"
	"github.com/stretchr/testify/assert"
)

func Test_Json(t *testing.T) {
	type ex struct {
		Name string `json:"n"`
		Age  int    `json:"i"`
	}

	w := httptest.NewRecorder()

	payload := ex{
		Name: "abc",
		Age:  20,
	}

	athena.Json(w, &payload, 204, nil)

	b, err := io.ReadAll(w.Body)
	if err != nil {
		assert.False(t, true, "Failed to read body for http body")
	}

	assert.Equal(t, `{"n":"abc","i":20}`, string(b))
	assert.Equal(t, 204, w.Result().StatusCode)
}
