package pubsub

import (
	"context"

	"github.com/IR-Digital-Token/x/messages"
)

type CallbackFn func(msg *messages.Message) error

type Pubsub interface {
	// Connect tries to create a connection to queue.
	Connect() error

	// Close function tries to close connection of queue.
	Close() error

	// Publish is used to insert a message to queue.
	Publish(ctx context.Context, topic string, message *messages.Message) error

	// Subscribe is used to receive messages from queue and handling them by "callback" function.
	Subscribe(ctx context.Context, topic string, callback CallbackFn) error
}
