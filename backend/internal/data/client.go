package data

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oseau/blockchain-explorer/ent"
)

type Client struct {
	db        *ent.Client
	ethClient *ethclient.Client
}

func NewClient() *Client {
	// TODO: move this to configs or secrets
	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/10eb704e03614a2d989085a1912b6344")
	if err != nil {
		log.Printf("create ethClient failed; err: %v\n", err)
	}
	return &Client{
		db:        ent.NewClient(),
		ethClient: ethClient,
	}
}
