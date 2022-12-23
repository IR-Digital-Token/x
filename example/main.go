package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"time"

	"github.com/IR-Digital-Token/x/chain"
	"github.com/IR-Digital-Token/x/chain/bindings/clipper"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/pubsub/gochannel"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	eth, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/qaVPRENeXs9vbDR8wYxKUmuIQhCakj0Y")
	if err != nil {
		log.Fatal(err)
	}

	q := gochannel.NewGoChannel(10)
	q.Subscribe(context.Background(), "clipper-kick", func(msg *messages.Message) error {
		log.Println("msg received")
		var v *clipper.ClipperKick
		err = json.Unmarshal(msg.Payload, &v)
		if err != nil {
			log.Println("error unmarshaling", err)
		}
		log.Println("kick in queue", v.Id, v.Usr)
		return nil
	})

	simepleKickHandler := clipper.NewKickHandler(common.HexToAddress("0xc67963a226eddd77B91aD8c421630A1b0AdFF270"), eth, simpleCallback())

	blockPtr := chain.NewFileBlockPointer(".", "eth.ptr", 16196820)
	blockPtr.Create()

	indexer := chain.NewIndexer(eth, blockPtr, 3)
	indexer.RegisterEventHandler(simepleKickHandler)
	indexer.RegisterAddress(common.HexToAddress("0xc67963a226eddd77B91aD8c421630A1b0AdFF270"))

	indexer.Init(time.Second * 10)
	for {
		err = indexer.Start()
		if err != nil {
			log.Println(err)
		}
	}
}

func simpleCallback() events.CallbackFn[clipper.ClipperKick] {
	return func(header types.Header, kick clipper.ClipperKick) error {
		fmt.Println("kick in simple", kick.Id, kick.Usr)
		return nil
	}
}
