package data

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

func (s Server) subscribeNewHeader() {
	newHeader := make(chan *types.Header)
	defer close(newHeader)

	sub, err := s.ethClient.SubscribeNewHead(context.Background(), newHeader)
	if err != nil {
		log.Printf("subscribeNewHeader failed; err: %v\n", err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case header := <-newHeader:
			log.Printf("get NewHeader: %v\n", header.Number)
			// *(s.blockNumber) = *header.Number
		case err := <-sub.Err():
			if err != nil {
				log.Printf("subscribeNewHeader err: %v\n", err)
			}
			return
		}
	}
}
