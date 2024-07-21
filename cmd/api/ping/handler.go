package ping

import (
	"net/http"
)

// HandlerV1 HTTP Handler of the endpoint /ping/v1
func HandlerV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	}
}
