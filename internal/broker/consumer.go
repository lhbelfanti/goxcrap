package broker

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

// NewConsumer creates a new pointer of RabbitMQMessageBroker for the Message Broker consumer
func NewConsumer(ctx context.Context, httpClient http.Client) (*RabbitMQMessageBroker, error) {
	consumer := &RabbitMQMessageBroker{
		httpClient: httpClient,
	}

	var err error
	consumer.conn, err = connect(ctx)
	if err != nil {
		return nil, err
	}
	ctx = log.With(ctx, log.Param("message_broker_consumer", consumer))

	consumer.channel, err = openChannel(ctx, consumer.conn)
	if err != nil {
		return nil, err
	}
	ctx = log.With(ctx, log.Param("message_broker_consumer", consumer))

	consumer.queue, err = declareQueue(ctx, consumer.channel)
	if err != nil {
		return nil, err
	}
	ctx = log.With(ctx, log.Param("message_broker_consumer", consumer))

	// Initialize the consumer
	consumer.messages, err = consumer.channel.ConsumeWithContext(
		ctx,
		consumer.queue.Name, // queue
		"",                  // consumer
		false,               // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, FailedToRegisterAConsumer
	}

	return consumer, nil
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

				// Ensure connection is still open before processing.
				// It could take many days to finish the processing, so the channel is probably closed.
				err = mb.ensureConnection(ctx)
				if err != nil {
					log.Error(ctx, err.Error())
				}

				err = msg.Ack(false)
				if err != nil {
					log.Error(ctx, err.Error())
				}
			}(msg)
		}
	}()
}
