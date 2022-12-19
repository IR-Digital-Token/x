package chain

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func HeadChannel(eth *ethclient.Client, interval time.Duration) (chan uint64, error) {
	headChan := make(chan uint64)
	go func() {
		t := time.NewTicker(interval)
		for {
			<-t.C
			block, err := eth.BlockNumber(context.Background())
			if err != nil {
				log.Println("cannot get head. retrying in a few seconds ...")
				time.Sleep(2 * time.Second)
				continue
			}
			headChan <- block
		}
	}()
	return headChan, nil
}
