package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IR-Digital-Token/x/chain"
	"github.com/IR-Digital-Token/x/chain/bindings/clipper"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/queue"
	"github.com/IR-Digital-Token/x/queue/gochannel"
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
	queueKickHandler := clipper.NewKickHandler(common.HexToAddress("0xc67963a226eddd77B91aD8c421630A1b0AdFF270"), eth, queueCallback(q))

	blockPtr := chain.NewFileBlockPointer(".", "eth.ptr", 16196820)
	blockPtr.Create()

	indexer := chain.NewIndexer(eth, blockPtr)
	indexer.RegisterEventHandler(simepleKickHandler)
	indexer.RegisterEventHandler(queueKickHandler)

	indexer.Init(time.Second * 10)
	indexer.Start()
}

func simpleCallback() events.CallbackFn[clipper.ClipperKick] {
	return func(kick clipper.ClipperKick) error {
		fmt.Println("kick in simple", kick.Id, kick.Usr)
		return nil
	}
}

func queueCallback(q queue.Q) events.CallbackFn[clipper.ClipperKick] {
	queue := q
	return func(event clipper.ClipperKick) error {
		payload, _ := json.Marshal(event)
		msg := messages.NewMessage(payload)
		err := queue.Publish(context.Background(), "clipper-kick", msg)
		if err != nil {
			log.Println("publish error", err)
		}
		log.Println("published")
		return nil
	}
}
