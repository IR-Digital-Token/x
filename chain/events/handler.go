package events

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type CallbackFn[T any] func(header types.Header, event T) error

type Handler interface {
	ID() string
	DecodeLog(log types.Log) (interface{}, error)
	HandleEvent(header types.Header, event interface{}) error
	DecodeAndHandle(header types.Header, log types.Log) error
}
