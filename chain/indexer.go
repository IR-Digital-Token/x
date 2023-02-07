package chain

import (
	"context"
	"sync"

	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/chain/transactions"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/panjf2000/ants/v2"

	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

type Test struct {
}
type Indexer struct {
	Head          chan uint64
	ptr           uint64
	batchSize     uint64
	blockInterval time.Duration
	pool          *ants.Pool
	eth           IEthereum
	blockPointer  BlockPointer
	logHandlers   map[string]events.Handler
	addresses     map[string]bool
	txWatchList   map[common.Hash]transactions.Handler
	mutex         *sync.Mutex
}

func NewIndexer(eth IEthereum, blockPointer BlockPointer, poolSize int) *Indexer {
	pool, err := ants.NewPool(poolSize)
	if err != nil {
		panic(err)
	}
	return &Indexer{
		eth:          eth,
		blockPointer: blockPointer,
		logHandlers:  make(map[string]events.Handler),
		pool:         pool,
		batchSize:    uint64(poolSize * 10),
		addresses:    make(map[string]bool),
		txWatchList:  make(map[common.Hash]transactions.Handler),
	}
}

func (w *Indexer) Init(blockInterval time.Duration) {
	ptr, err := w.blockPointer.Read()
	if err != nil {
		panic(err)
	}
	w.ptr = ptr

	head, err := HeadChannel(w.eth, blockInterval)
	if err != nil {
		panic(err)
	}
	w.Head = head
	w.blockInterval = blockInterval
}

func (w *Indexer) Start() error {
	head := <-w.Head
	for w.ptr < head {
		err := w.loop(w.ptr, min(w.ptr+w.batchSize, head))
		if err != nil {
			return err
		}

		diff := head - w.ptr
		if diff < w.batchSize {
			w.ptr += diff
		} else {
			w.ptr += w.batchSize
		}
		err = w.blockPointer.Update(w.ptr)
		if err != nil {
			return err
		}
	}
	return nil
}

func min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func (w *Indexer) loop(from, to uint64) error {
	ch := make(chan error)
	done := make(chan struct{})
	parent := context.Background()
	go func() {
		wg := &sync.WaitGroup{}
		for i := from; i < to; i++ {
			wg.Add(1)
			j := i
			err := w.pool.Submit(func() {
				ctx, cancel := context.WithTimeout(parent, w.blockInterval)
				defer cancel()
				err := w.processBlock(ctx, big.NewInt(int64(j)))
				if err != nil {
					ch <- err
					return
				}
				wg.Done()
			})
			if err != nil {
				ch <- err
				return
			}
		}
		wg.Wait()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return nil
	case err := <-ch:
		return err
	}
}

func (w *Indexer) processBlock(ctx context.Context, number *big.Int) error {
	block, err := w.eth.BlockByNumber(ctx, number)
	if err != nil {
		return err
	}

	logs, err := w.eth.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: block.Number(),
		ToBlock:   block.Number(),
	})
	if err != nil {
		return err
	}
	err = w.processTransactions(*block.Header(), w.filterTxHash(block.Transactions()))
	if err != nil {
		return err
	}
	return w.processLogs(*block.Header(), w.filterLogs(logs))
}

func (w *Indexer) processTransactions(header types.Header, txList types.Transactions) error {
	for _, tx := range txList {
		txHash := tx.Hash()
		txRecipt, err := w.eth.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return err
		}

		handler, ok := w.txWatchList[txHash]
		if !ok {
			continue
		}
		err = handler.HandleTransaction(header, txRecipt)
		if err != nil {
			return err
		}
		w.UnWatchTx(handler)
	}
	return nil
}

func (w *Indexer) processLogs(header types.Header, logs []types.Log) error {
	for _, l := range logs {
		if len(l.Topics) == 0 {
			continue
		}
		handler, ok := w.logHandlers[l.Topics[0].String()]
		if !ok {
			continue
		}
		err := handler.DecodeAndHandle(header, l)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Indexer) filterLogs(logs []types.Log) []types.Log {
	var res []types.Log
	for _, l := range logs {
		_, ok := w.addresses[l.Address.String()]

		if ok {
			res = append(res, l)
		}
	}
	return res
}
func (w *Indexer) filterTxHash(transactions types.Transactions) types.Transactions {
	var res types.Transactions
	for _, tx := range transactions {
		txHash := tx.Hash()
		_, ok := w.txWatchList[txHash]
		if ok {
			res = append(res, tx)
		}
	}
	return res
}

func (w *Indexer) RegisterEventHandler(handler events.Handler) {
	w.logHandlers[handler.ID()] = handler
}
func (w *Indexer) RegisterEventHandlersDon(handlers ...events.Handler) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for _, handler := range handlers {
		w.logHandlers[handler.ID()] = handler
	}
}

func (w *Indexer) RegisterAddress(addr common.Address) {
	w.addresses[addr.String()] = true
}

func (w *Indexer) WatchTx(handler transactions.Handler) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.txWatchList[handler.ID()] = handler
}

func (w *Indexer) UnWatchTx(txHandler transactions.Handler) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	delete(w.txWatchList, txHandler.ID())
}

func (w *Indexer) Ptr() uint64 {
	return w.ptr
}
