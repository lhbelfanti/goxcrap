package ping_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/ping"
)

func TestRunHandlerV1_success(t *testing.T) {
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/ping/v1", strings.NewReader(""))

	handlerV1 := ping.HandlerV1()

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
