package chain

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestRegisterAddress(t *testing.T) {
	account := common.HexToAddress("0x28A86dd85bCc6773942B923Ff988AF5f85398115")
	var eth *ethclient.Client
	var blockInterval BlockPointer
	indexer := NewIndexer(eth, blockInterval, 2)
	indexer.RegisterAddress(account)
	if indexer.addresses[account.String()] {
		t.Logf("expect %s is equal to %t in address and it's %t", account, true, true)
	} else {
		t.Errorf("expect %s is equal to %t in address and it's %t", account, true, false)
	}
}

// func TestWatchTx(t *testing.T) {
// 	txHash := common.HexToHash("0xa3e4704298180c945838738728594b4d5da36d6c51bfc946ba31317646b61646")
// 	var eth *ethclient.Client
// 	var blockInterval BlockPointer
// 	indexer := NewIndexer(eth, blockInterval, 2)
// 	var txHandler transactions.Handler
// 	t.Log(txHandler.ID())
// 	indexer.WatchTx(txHandler)
// 	handler := indexer.txWatchList[txHandler.ID()]
// 	t.Log(txHandler, txHash)
// 	if reflect.TypeOf(handler).Kind() == transactions.CallbackFn {
// 		t.Logf("expect %s is equal to %t in address and it's %t", account, true, true)
// 	} else {
// 		t.Errorf("expect %s is equal to %t in address and it's %t", account, true, false)
// 	}
// }
