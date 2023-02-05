package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
)

type Chain interface {
	BlockNumber(ctx context.Context) (uint64, error)
	ChainID(ctx context.Context) (*big.Int, error)
}

type IEthereum interface {
	ethereum.ChainReader
	ethereum.LogFilterer
	ethereum.ChainStateReader
	ethereum.ChainSyncReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.PendingContractCaller
	ethereum.PendingStateEventer
	ethereum.PendingStateReader
	ethereum.Subscription
	ethereum.TransactionReader
	ethereum.TransactionSender
	Chain
}
