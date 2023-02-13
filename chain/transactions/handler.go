package transactions

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CallbackFn func(header types.Header, recipt *types.Receipt) error
type Handler interface {
	ID() common.Hash
	HandleTransaction() func(header types.Header, recipt *types.Receipt) error
}
