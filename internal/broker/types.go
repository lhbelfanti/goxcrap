package broker

import (
	"context"
	"fmt"
	"os"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/http"
)

const rabbitmqURL = "amqp://%s:%s@rabbitmq:%s/"

type (
	// MessageBroker defines the necessary methods for a message broker implementation
	MessageBroker interface {
		// EnqueueMessage enqueues a new message in the messages queue
		EnqueueMessage(ctx context.Context, body string) error

		// InitMessageConsumerWithEndpoint initializes the goroutine in charge of the messages consumption
		// It calls the endpoint passed by parameter to process the message
		// Choose this method OR InitMessageConsumerWithFunction to consume the messages
		// Using both at the same time will cause a race condition between them to process the messages
		InitMessageConsumerWithEndpoint(concurrentMessages int, processorEndpoint string)

		// InitMessageConsumerWithFunction initializes the goroutine in charge of the messages consumption
		// It executes the function passed by parameter to process de message
		// Choose this method OR InitMessageConsumerWithEndpoint to consume the messages.
		// Using both at the same time will cause a race condition between them to process the messages
		InitMessageConsumerWithFunction(concurrentMessages int, processorFunc ProcessorFunction)
	}

	// ProcessorFunction function in charge of processing the messages of the MessageBroker
	// It is part of the method of consumption: InitMessageConsumerWithFunction
	ProcessorFunction func(ctx context.Context, body []byte) error

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
	port := os.Getenv("RABBITMQ_PORT")

	return fmt.Sprintf(rabbitmqURL, user, pass, port)
}
