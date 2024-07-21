package broker

import "errors"

var (
	FailedToInitializeRabbitMQ = errors.New("failed to initialize RabbitMQ")
	FailedToOpenAChannel       = errors.New("failed to open a channel")
	FailedToDeclareAQueue      = errors.New("failed to declare a queue")
	FailedToRegisterAConsumer  = errors.New("failed to register a consumer")
	FailedToPublishAPublishing = errors.New("failed to publish a publishing")
)
