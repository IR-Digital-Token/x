package chain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
func sendNewTransaction(t *testing.T) (SimulatedEthereum, *types.Receipt, *types.Block, common.Hash) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei
	fmt.Printf("t: %v\n", balance)
	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	blockGasLimit := uint64(4712388)
	client := NewSimulatedEthereum(backends.NewSimulatedBackend(genesisAlloc, blockGasLimit))

	// fromAddress := auth.From
	var nonce uint64
	nonce = 0
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	value := big.NewInt(0)    // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(nil), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}

	client.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	block, err := client.BlockByHash(context.Background(), receipt.BlockHash)
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	return client, receipt, block, signedTx.Hash()
}
func TestWatchTx(t *testing.T) {
	// var cnt int
	eth, recipt, block, txHash := sendNewTransaction(t)
	fmt.Print(eth, recipt, block, txHash)
	// var blockInterval BlockPointer
	// indexer := NewIndexer(eth, blockInterval, 2)
	// mockTxHandler := &repomocks.Handler{}
	// mockTxHandler.On("ID").Return(txHash).Once()
	// mockTxHandler.On("HandleTransaction").Return(func(header types.Header, recipt *types.Receipt) error {
	// 	cnt += 1
	// 	return nil
	// }).Once()
	// indexer.WatchTx(mockTxHandler)
	// indexer.txWatchList[txHash].HandleTransaction(*block.Header(), recipt)
	// assert.Equal(t, 1, cnt)
}
