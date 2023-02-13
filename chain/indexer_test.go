package chain

import (
	"context"
	"log"
	"math/big"
	"testing"

	repomocks "github.com/IR-Digital-Token/x/chain/transactions/mocks"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei
	address := auth.From

	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	blockGasLimit := uint64(4712388)
	client := NewSimulatedEthereum(backends.NewSimulatedBackend(genesisAlloc, blockGasLimit))

	client.Commit()
	var nonce uint64
	nonce = 0
	value := big.NewInt(0)    // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), privateKey)
	if err != nil {
		t.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	var blockInterval BlockPointer
	var cnt int
	indexer := NewIndexer(client, blockInterval, 2)
	mockTxHandler := repomocks.NewHandler(t)
	mockTxHandler.On("ID").Return(signedTx.Hash()).On("HandleTransaction").Return(func(header types.Header, recipt *types.Receipt) error {
		cnt += 1
		return nil
	})
	indexer.WatchTx(mockTxHandler)

	client.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	block, err := client.BlockByHash(context.Background(), receipt.BlockHash)
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}
	checkHandler := indexer.txWatchList[signedTx.Hash()]

	mockTxHandler.HandleTransaction()(*block.Header(), receipt)
	assert.Equal(t, signedTx.Hash(), checkHandler.ID(), "The two words should be the same.")
	assert.Equal(t, cnt, 1, "callback function must increase the counter after geting the transaction hash from block tx list")
}
