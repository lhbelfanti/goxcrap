package broker

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

// NewMessageBroker creates a new pointer of RabbitMQMessageBroker
func NewMessageBroker(ctx context.Context, httpClient http.Client) (*RabbitMQMessageBroker, error) {
	messageBroker := &RabbitMQMessageBroker{
		httpClient: httpClient,
	}
	var err error

	// Connect to RabbitMQ
	url := resolveRabbitmqURL()
	ctx = log.With(ctx, log.Param("rabbitmq_url", url))
	messageBroker.conn, err = amqp091.Dial(url)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, FailedToInitializeRabbitMQ
	}
	ctx = log.With(ctx, log.Param("message_broker", messageBroker))

	// Create a channel
	messageBroker.channel, err = messageBroker.conn.Channel()
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, FailedToOpenAChannel
	}
	ctx = log.With(ctx, log.Param("message_broker", messageBroker))

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
		log.Error(ctx, err.Error())
		return nil, FailedToDeclareAQueue
	}
	ctx = log.With(ctx, log.Param("message_broker", messageBroker))

	// Initialize the consumer
	messageBroker.messages, err = messageBroker.channel.ConsumeWithContext(
		ctx,
		messageBroker.queue.Name, // queue
		"",                       // consumer
		false,                    // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, FailedToRegisterAConsumer
	}

	return messageBroker, nil
}

// EnqueueMessage enqueues a new message in the RabbitMQMessageBroker.queue
func (mb *RabbitMQMessageBroker) EnqueueMessage(ctx context.Context, body string) error {
	ctx = log.With(ctx, log.Param("body", body))

	publishing := amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         []byte(body),
	}

	err := mb.channel.PublishWithContext(
		ctx,
		"",            // exchange
		mb.queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		publishing,
	)
	if err != nil {
		log.Error(ctx, err.Error())
		return FailedToPublishAPublishing
	}

	return nil
}

// InitMessageConsumerWithEndpoint is a goroutine that is constantly processing the enqueued messages
// It receives two params:
// concurrentMessages: defines the amount of messages that can be processed in parallel
// processorEndpoint: defines the endpoint that is called when a message from the queue is consumed.
// That endpoint is in charge of processing the enqueued messages
func (mb *RabbitMQMessageBroker) InitMessageConsumerWithEndpoint(concurrentMessages int, processorEndpoint string) {
	workerChan := make(chan struct{}, concurrentMessages)

	go func() {
		for msg := range mb.messages {
			workerChan <- struct{}{}
			go func(msg amqp091.Delivery) {
				defer func() {
					<-workerChan
				}()

				ctx := context.Background()
				resp, err := mb.httpClient.NewRequest(ctx, "POST", processorEndpoint, msg.Body)
				if err != nil {
					log.Error(ctx, err.Error())
				}
				ctx = log.With(ctx, log.Param("msg_body", string(msg.Body)))

				log.Info(ctx, fmt.Sprintf("Message consumer endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

				err = msg.Ack(false)
				if err != nil {
					log.Error(ctx, err.Error())
				}
			}(msg)
		}
	}()
}

// InitMessageConsumerWithFunction is a goroutine that is constantly processing the enqueued messages
// It receives two params:
// concurrentMessages: defines the amount of messages that can be processed in parallel
// processorFunc: defines the function that is called when a message from the queue is consumed.
// That function is in charge of processing the enqueued messages
func (mb *RabbitMQMessageBroker) InitMessageConsumerWithFunction(concurrentMessages int, processorFunc ProcessorFunction) {
	workerChan := make(chan struct{}, concurrentMessages)

	go func() {
		for msg := range mb.messages {
			workerChan <- struct{}{}
			go func(msg amqp091.Delivery) {
				defer func() {
					<-workerChan
				}()

				ctx := context.Background()
				err := processorFunc(ctx, msg.Body)
				if err != nil {
					log.Error(ctx, err.Error())
				}
				ctx = log.With(ctx, log.Param("msg_body", string(msg.Body)))

				log.Info(ctx, "Message consumer function called")

				err = msg.Ack(false)
				if err != nil {
					log.Error(ctx, err.Error())
				}
			}(msg)
		}
	}()
}
