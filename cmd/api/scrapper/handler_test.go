package scrapper_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/scrapper"
)

func TestExecuteHandlerV1_success(t *testing.T) {
	mockExecute := scrapper.MockExecute(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/execute-scrapper/v1", strings.NewReader(""))

	handlerV1 := scrapper.ExecuteHandlerV1(mockExecute)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsError(t *testing.T) {
	mockExecute := scrapper.MockExecute(errors.New("execute scrapper failed"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/execute-scrapper/v1", strings.NewReader(""))

	handlerV1 := scrapper.ExecuteHandlerV1(mockExecute)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}
