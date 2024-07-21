package broker

import "github.com/stretchr/testify/mock"

// MockMessageBroker is a mock implementation of RabbitMQMessageBroker
type MockMessageBroker struct {
	mock.Mock
}

func (m *MockMessageBroker) EnqueueMessage(body string) error {
	args := m.Called(body)
	return args.Error(0)
}

func (m *MockMessageBroker) InitMessageConsumer(int, string) {
	return
}
