package broker

import (
	"context"
	"fmt"
	"os"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/http"
)

const rabbitmqURL = "amqp://%s:%s@rabbitmq:5672/"

type (
	// MessageBroker defines the necessary methods for a message broker implementation
	MessageBroker interface {
		// EnqueueMessage enqueues a new message in the messages queue
		EnqueueMessage(ctx context.Context, body string) error

		// InitMessageConsumer initializes the goroutine in charge of the messages consumption
		InitMessageConsumer(ctx context.Context, concurrentMessages int, processorEndpoint string)
	}

	// RabbitMQMessageBroker contains all the necessary variables for a message broker: the connection, the channel, the queue, and a chan of messages
	RabbitMQMessageBroker struct {
		conn       *amqp091.Connection
		channel    *amqp091.Channel
		queue      amqp091.Queue
		messages   <-chan amqp091.Delivery
		httpClient http.Client
	}
)

func resolveRabbitmqURL() string {
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")

	return fmt.Sprintf(rabbitmqURL, user, pass)
}
