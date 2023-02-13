package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/ethereum/go-ethereum"
)

type Blockchain interface {
	BlockNumber(ctx context.Context) (uint64, error)
	ChainID(ctx context.Context) (*big.Int, error)
}

type SimulatedBlockchain interface {
	Commit() common.Hash
	Rollback()
	Fork(ctx context.Context, parent common.Hash) error
}

type Ethereum interface {
	ethereum.ChainReader
	ethereum.LogFilterer
	ethereum.ChainStateReader
	ethereum.ChainSyncReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.PendingContractCaller
	ethereum.PendingStateReader
	ethereum.TransactionReader
	ethereum.TransactionSender
	Blockchain
}
