package chain

import (
	"testing"

	repomocks "github.com/IR-Digital-Token/x/chain/transactions/mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAddress(t *testing.T) {
	account := common.HexToAddress("0x28A86dd85bCc6773942B923Ff988AF5f85398115")
	var eth SimulatedEthereum
	var blockInterval BlockPointer
	indexer := NewIndexer(eth, blockInterval, 2)
	indexer.RegisterAddress(account)
	if indexer.addresses[account.String()] {
		t.Logf("expect %s is equal to %t in address and it's %t", account, true, true)
	} else {
		t.Errorf("expect %s is equal to %t in address and it's %t", account, true, false)
	}
}

func TestWatchTx(t *testing.T) {
	var cnt int
	txHash := common.HexToHash("0xa3e4704298180c945838738728594b4d5da36d6c51bfc946ba31317646b61646")
	var eth SimulatedEthereum
	var blockInterval BlockPointer
	indexer := NewIndexer(eth, blockInterval, 2)
	mockTxHandler := &repomocks.Handler{}
	mockTxHandler.On("ID").Return(txHash).Once()
	mockTxHandler.On("HandleTransaction").Return(func(header types.Header, recipt *types.Receipt) error {
		cnt += 1
		return nil
	}).Once()
	indexer.WatchTx(mockTxHandler)
	indexer.txWatchList[txHash].HandleTransaction(types.Header{}, &types.Receipt{})
	assert.Equal(t, 1, cnt)
}
