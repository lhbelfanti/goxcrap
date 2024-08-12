package broker

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockMessageBroker is a mock implementation of RabbitMQMessageBroker
type MockMessageBroker struct {
	mock.Mock
}

func (m *MockMessageBroker) EnqueueMessage(ctx context.Context, body string) error {
	args := m.Called(body)
	return args.Error(0)
}

func (m *MockMessageBroker) InitMessageConsumer(ctx context.Context, concurrentMessages int, processorEndpoint string) {
	return
}
