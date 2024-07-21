package broker

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type (
	// MessageBroker defines the necessary methods for a message broker implementation
	MessageBroker interface {
		// EnqueueMessage enqueues a new message in the messages queue
		EnqueueMessage(body string) error

		// InitMessageConsumer initializes the goroutine in charge of the messages consumption
		InitMessageConsumer(concurrentMessages int, processorEndpoint string)
	}

	// RabbitMQMessageBroker contains all the necessary variables for a message broker: the connection, the channel, the queue, and a chan of messages
	RabbitMQMessageBroker struct {
		conn     *amqp091.Connection
		channel  *amqp091.Channel
		queue    amqp091.Queue
		messages <-chan amqp091.Delivery
	}
)

const rabbitmqURL = "amqp://%s:%s@rabbitmq:5672/"

func resolveRabbitmqURL() string {
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")

	url := fmt.Sprintf(rabbitmqURL, user, pass)

	slog.Info(url)
	return url
}
