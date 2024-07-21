package criteria_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
)

func TestEnqueueHandlerV1_success(t *testing.T) {
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything).Return(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/enqueue-criteria/v1", bytes.NewReader(mockBody))

	handlerV1 := criteria.EnqueueHandlerV1(mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failsWhenTheBodyIsInvalid(t *testing.T) {
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything).Return(nil)
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/enqueue-criteria/v1", bytes.NewReader(mockBody))

	handlerV1 := criteria.EnqueueHandlerV1(mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failsWhenEnqueueMessageThrowsError(t *testing.T) {
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything).Return(errors.New("error while executing EnqueueMessage"))
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/enqueue-criteria/v1", bytes.NewReader(mockBody))

	handlerV1 := criteria.EnqueueHandlerV1(mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}
