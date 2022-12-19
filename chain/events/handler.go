package events

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type CallbackFn[T any] func(event T) error

type Handler interface {
	Signature() string
	DecodeLog(log types.Log) (interface{}, error)
	HandleEvent(event interface{}) error
	DecodeAndHandle(log types.Log) error
}
