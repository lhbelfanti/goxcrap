package broker

import (
	"context"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/log"
)

// openChannel creates a new Message Broker Channel for a given Message Broker Connection and returns it
func openChannel(ctx context.Context, connection *amqp091.Connection) (*amqp091.Channel, error) {
	channel, err := connection.Channel()
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, FailedToOpenAChannel
	}

	return channel, nil
}
