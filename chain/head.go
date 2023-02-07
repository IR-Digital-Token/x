package chain

import (
	"context"
	"log"
	"time"
)

func HeadChannel(eth Ethereum, interval time.Duration) (chan uint64, error) {
	headChan := make(chan uint64)
	go func() {
		t := time.NewTicker(interval)
		for {
			<-t.C
			block, err := eth.BlockNumber(context.Background())
			if err != nil {
				log.Println("error getting head.", err)
				time.Sleep(2 * time.Second)
				continue
			}
			headChan <- block
		}
	}()
	return headChan, nil
}
