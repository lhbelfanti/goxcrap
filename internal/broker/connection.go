package broker

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"

	"goxcrap/internal/log"
)

// connect establish a new connection with the Message Broker and returns it
func connect(ctx context.Context) (*amqp091.Connection, error) {
	url := resolveRabbitmqURL()
	ctx = log.With(ctx, log.Param("rabbitmq_url", url))
	connection, err := amqp091.DialConfig(url,
		amqp091.Config{
			Heartbeat: 30 * time.Second,
		},
	)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, FailedToInitializeRabbitMQ
	}

	return connection, nil
}

func resolveRabbitmqURL() string {
	const rabbitmqURL = "amqp://%s:%s@rabbitmq:%s/"

	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")
	port := os.Getenv("RABBITMQ_PORT")

	return fmt.Sprintf(rabbitmqURL, user, pass, port)
}

func (mb *RabbitMQMessageBroker) CloseConnection() {
	if mb.conn != nil {
		_ = mb.conn.Close() // Close the old connection if any
	}
}

// reconnect Reconnects to RabbitMQ.
// RabbitMQ's connections and channels can be closed due to inactivity, errors, or network issues
func (mb *RabbitMQMessageBroker) reconnect() error {
	mb.CloseConnection()

	var err error
	url := resolveRabbitmqURL()
	mb.conn, err = amqp091.Dial(url)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	mb.channel, err = mb.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	return nil
}

// ensureConnection Reconnect if the connection is lost
func (mb *RabbitMQMessageBroker) ensureConnection(ctx context.Context) error {
	if mb.conn == nil || mb.conn.IsClosed() {
		log.Info(ctx, "Reconnecting to RabbitMQ")
		return mb.reconnect()
	}

	return nil
}
