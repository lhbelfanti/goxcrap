package ping

import (
	"encoding/json"
	"net/http"

	"goxcrap/internal/http/response"
)

// HandlerV1 HTTP Handler of the endpoint /ping/v1
func HandlerV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := response.DTO{Code: http.StatusOK, Message: "pong"}
		_ = json.NewEncoder(w).Encode(resp)
	}
}
