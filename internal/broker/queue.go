package broker

import (
	"context"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/log"
)

// declareQueue declares a new Message Broker Queue and returns it
func declareQueue(ctx context.Context, channel *amqp091.Channel) (amqp091.Queue, error) {
	queue, err := channel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Error(ctx, err.Error())
		return amqp091.Queue{}, FailedToDeclareAQueue
	}

	return queue, nil
}
