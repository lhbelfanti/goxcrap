package broker

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/rabbitmq/amqp091-go"
)

// NewMessageBroker creates a new pointer of RabbitMQMessageBroker
func NewMessageBroker() (*RabbitMQMessageBroker, error) {
	messageBroker := &RabbitMQMessageBroker{}
	var err error

	// Connect to RabbitMQ
	messageBroker.conn, err = amqp091.Dial(resolveRabbitmqURL())
	if err != nil {
		slog.Error(err.Error())
		return nil, FailedToInitializeRabbitMQ
	}

	// Create a channel
	messageBroker.channel, err = messageBroker.conn.Channel()
	if err != nil {
		slog.Error(err.Error())
		return nil, FailedToOpenAChannel
	}

	// Declare a queue
	messageBroker.queue, err = messageBroker.channel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		slog.Error(err.Error())
		return nil, FailedToDeclareAQueue
	}

	// Initialize the consumer
	messageBroker.messages, err = messageBroker.channel.Consume(
		messageBroker.queue.Name, // queue
		"",                       // consumer
		false,                    // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		slog.Error(err.Error())
		return nil, FailedToRegisterAConsumer
	}

	return messageBroker, nil
}

// EnqueueMessage enqueues a new message in the RabbitMQMessageBroker.queue
func (mb *RabbitMQMessageBroker) EnqueueMessage(body string) error {
	publishing := amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         []byte(body),
	}

	err := mb.channel.Publish(
		"",            // exchange
		mb.queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		publishing,
	)
	if err != nil {
		slog.Error(err.Error())
		return FailedToPublishAPublishing
	}

	return nil
}

// InitMessageConsumer is a goroutine that is constantly processing the enqueued messages
// It receives two params:
// concurrentMessages: defines the amount of messages that can be processed in parallel
// processorEndpoint: defines the endpoint that is called when a message from the queue is consumed. That endpoint is in charge of processing the enqueued messages
func (mb *RabbitMQMessageBroker) InitMessageConsumer(concurrentMessages int, processorEndpoint string) {
	workerChan := make(chan struct{}, concurrentMessages)

	go func() {
		for msg := range mb.messages {
			workerChan <- struct{}{}
			go func(msg amqp091.Delivery) {
				defer func() {
					<-workerChan
				}()
				// TODO replace callEndpoint by the http client of the internal/http
				callEndpoint(processorEndpoint, string(msg.Body))
				err := msg.Ack(false)
				if err != nil {
					slog.Error(err.Error())
				}
			}(msg)
		}
	}()
}

// callEndpoint calls the POST endpoint of this application passed by parameter. It also receives its body as string
func callEndpoint(endpoint, body string) {
	url := fmt.Sprintf("http://localhost:8091%s", endpoint)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(resp.Body)

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info(string(response))
}
