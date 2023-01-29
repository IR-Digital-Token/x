package transactions

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Handler interface {
	ID() common.Hash
	HandleTransaction(header types.Header, recipt *types.Receipt) error
}
