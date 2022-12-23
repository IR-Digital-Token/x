package chain

import (
	"context"

	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum"

	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Indexer struct {
	Head chan uint64

	ptr uint64

	eth                 *ethclient.Client
	blockPositionHolder BlockPointer
	logHandlers         map[string]events.Handler
}

func NewIndexer(eth *ethclient.Client, blockPointer BlockPointer) *Indexer {
	return &Indexer{
		eth:                 eth,
		blockPositionHolder: blockPointer,
		logHandlers:         make(map[string]events.Handler),
	}
}

func (w *Indexer) Init(blockInterval time.Duration) {
	ptr, err := w.blockPositionHolder.Read()
	if err != nil {
		panic(err)
	}
	w.ptr = ptr

	head, err := HeadChannel(w.eth, blockInterval)
	if err != nil {
		panic(err)
	}
	w.Head = head
}

func (w *Indexer) Start() error {
	head := <-w.Head

	for w.ptr <= head {
		block, err := w.eth.BlockByNumber(context.Background(), big.NewInt(int64(w.ptr)))
		if err != nil {
			return err
		}

		logs, err := w.eth.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: block.Number(),
			ToBlock:   block.Number(),
		})
		if err != nil {
			return err
		}

		err = w.processLogs(logs)
		if err != nil {
			return err
		}

		err = w.blockPositionHolder.Update(w.ptr)
		if err != nil {
			return err
		}

		w.ptr++
	}

	return nil
}

func (w *Indexer) processLogs(logs []types.Log) error {
	for _, l := range logs {
		if len(l.Topics) == 0 {
			continue
		}
		handler, ok := w.logHandlers[l.Topics[0].String()]
		if !ok {
			continue
		}
		err := handler.DecodeAndHandle(l)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Indexer) RegisterEventHandler(handler events.Handler) {
	w.logHandlers[handler.ID()] = handler
}

func (w *Indexer) Ptr() uint64 {
	return w.ptr
}
