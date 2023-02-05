package chain

import (
	"context"
	"log"
	"math/big"
	"testing"

	repomocks "github.com/IR-Digital-Token/x/chain/transactions/mocks"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAddress(t *testing.T) {
	account := common.HexToAddress("0x28A86dd85bCc6773942B923Ff988AF5f85398115")
	var eth *ethereum.ChainReader
	var blockInterval BlockPointer
	indexer := NewIndexer(eth, blockInterval, 2)
	indexer.RegisterAddress(account)
	if indexer.addresses[account.String()] {
		t.Logf("expect %s is equal to %t in address and it's %t", account, true, true)
	} else {
		t.Errorf("expect %s is equal to %t in address and it's %t", account, true, false)
	}
}

func newSimulatedBackend() (backends.SimulatedBackend, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Panic(err)
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
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

}
func TestWatchTx(t *testing.T) {
	var cnt int
	txHash := common.HexToHash("0xa3e4704298180c945838738728594b4d5da36d6c51bfc946ba31317646b61646")
	var eth *ethclient.Client
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
