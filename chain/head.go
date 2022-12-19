package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func HeadChannel(c *ethclient.Client, interval time.Duration) (chan uint64, error) {
	headChan := make(chan uint64)
	go func() {
		t := time.NewTicker(interval)
		for {
			<-t.C
			block, err := c.BlockNumber(context.Background())
			if err != nil {
				log.Println("cannot get head. retry in 2s ...")
				time.Sleep(2 * time.Second)
				continue
			}
			headChan <- block
		}
	}()
	return headChan, nil
}
