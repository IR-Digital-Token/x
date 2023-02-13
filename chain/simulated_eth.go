package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimulatedEthereum interface {
	Ethereum
	SimulatedBlockchain
}

type simulatedEthereum struct {
	backEnd *backends.SimulatedBackend
}

func (s *simulatedEthereum) Commit() common.Hash {
	return s.backEnd.Commit()
}

func (s *simulatedEthereum) Rollback() {
	s.backEnd.Rollback()
}

func (s *simulatedEthereum) Fork(ctx context.Context, parent common.Hash) error {
	return s.Fork(ctx, parent)
}

func (s *simulatedEthereum) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return s.backEnd.BlockByHash(ctx, hash)
}

func (s *simulatedEthereum) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return s.BlockByNumber(ctx, number)
}

func (s *simulatedEthereum) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return s.HeaderByHash(ctx, hash)
}

func (s *simulatedEthereum) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return s.HeaderByNumber(ctx, number)
}

func (s *simulatedEthereum) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return s.TransactionCount(ctx, blockHash)
}

func (s *simulatedEthereum) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	return s.TransactionInBlock(ctx, blockHash, index)
}

func (s *simulatedEthereum) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return s.SubscribeNewHead(ctx, ch)
}

func (s *simulatedEthereum) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return s.backEnd.FilterLogs(ctx, q)
}

func (s *simulatedEthereum) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return s.SubscribeFilterLogs(ctx, q, ch)
}

func (s *simulatedEthereum) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return s.backEnd.BalanceAt(ctx, account, blockNumber)
}

func (s *simulatedEthereum) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return s.backEnd.StorageAt(ctx, account, key, blockNumber)
}

func (s *simulatedEthereum) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return s.CodeAt(ctx, account, blockNumber)
}

func (s *simulatedEthereum) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return s.backEnd.NonceAt(ctx, account, blockNumber)
}

func (s *simulatedEthereum) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return s.SyncProgress(ctx)
}

func (s *simulatedEthereum) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return s.CallContract(ctx, call, blockNumber)
}

func (s *simulatedEthereum) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return s.EstimateGas(ctx, call)
}

func (s *simulatedEthereum) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return s.backEnd.SuggestGasPrice(ctx)
}

func (s *simulatedEthereum) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return s.PendingCallContract(ctx, call)
}

func (s *simulatedEthereum) SubscribePendingTransactions(ctx context.Context, ch chan<- *types.Transaction) (ethereum.Subscription, error) {
	return s.SubscribePendingTransactions(ctx, ch)
}

func (s *simulatedEthereum) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	panic("no implementation")
}

func (s *simulatedEthereum) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	panic("no implementation")
}

func (s *simulatedEthereum) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return s.backEnd.PendingCodeAt(ctx, account)
}

func (s *simulatedEthereum) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return s.PendingNonceAt(ctx, account)
}

func (s *simulatedEthereum) PendingTransactionCount(ctx context.Context) (uint, error) {
	panic("no implementation")
}

func (s *simulatedEthereum) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return s.backEnd.TransactionByHash(ctx, txHash)
}

func (s *simulatedEthereum) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return s.backEnd.TransactionReceipt(ctx, txHash)
}

func (s *simulatedEthereum) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return s.backEnd.SendTransaction(ctx, tx)
}

func (s *simulatedEthereum) BlockNumber(ctx context.Context) (uint64, error) {
	return s.backEnd.Blockchain().CurrentBlock().NumberU64(), nil
}

func (s *simulatedEthereum) ChainID(ctx context.Context) (*big.Int, error) {
	return s.backEnd.Blockchain().Config().ChainID, nil
}

func NewSimulatedEthereum(backend *backends.SimulatedBackend) SimulatedEthereum {
	return &simulatedEthereum{backEnd: backend}
}
