package broker

import (
	"context"
	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

// NewProducer creates a new pointer of RabbitMQMessageBroker for the Message Broker producer
func NewProducer(ctx context.Context, httpClient http.Client) (*RabbitMQMessageBroker, error) {
	producer := &RabbitMQMessageBroker{
		httpClient: httpClient,
	}

	var err error
	producer.conn, err = connect(ctx)
	if err != nil {
		return nil, err
	}
	ctx = log.With(ctx, log.Param("message_broker_producer", producer))

	producer.channel, err = openChannel(ctx, producer.conn)
	if err != nil {
		return nil, err
	}
	ctx = log.With(ctx, log.Param("message_broker_producer", producer))

	producer.queue, err = declareQueue(ctx, producer.channel)
	if err != nil {
		return nil, err
	}
	ctx = log.With(ctx, log.Param("message_broker_producer", producer))

	return producer, nil
}

// EnqueueMessage enqueues a new message in the RabbitMQMessageBroker.queue
func (mb *RabbitMQMessageBroker) EnqueueMessage(ctx context.Context, body string) error {
	ctx = log.With(ctx, log.Param("body", body))

	err := mb.channel.PublishWithContext(
		ctx,
		"",            // exchange
		mb.queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		},
	)
	if err != nil {
		log.Error(ctx, err.Error())
		return FailedToPublishAPublishing
	}

	return nil
}
