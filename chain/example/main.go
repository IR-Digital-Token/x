package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IR-Digital-Token/x/chain"
	clipperBinding "github.com/IR-Digital-Token/x/chain/bindings/clipper"
	"github.com/IR-Digital-Token/x/chain/events/clipper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	eth, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/qaVPRENeXs9vbDR8wYxKUmuIQhCakj0Y")
	if err != nil {
		log.Fatal(err)
	}

	kickHandler := clipper.NewKickHandler(common.HexToAddress("0xc67963a226eddd77B91aD8c421630A1b0AdFF270"), eth, func(kick clipperBinding.ClipperKick) error {
		fmt.Println("kick", kick.Id, kick.Usr)
		os.Exit(0)
		return nil
	})

	blockPtr := chain.NewFileBlockPointer(".", "eth.ptr", 16196820)
	blockPtr.Create()

	indexer := chain.NewIndexer(eth, blockPtr)
	indexer.RegisterEventHandler(kickHandler)

	indexer.Init(time.Second * 10)
	indexer.Start()
}
