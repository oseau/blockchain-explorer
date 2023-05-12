package data

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oseau/blockchain-explorer/ent"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	db            *ent.Client
	ethClient     *ethclient.Client
	moralisConfig *moralisConfig
}

type moralisConfig struct {
	url string
	key string // TODO: protect this
}

func NewServer() *Server {
	// TODO: move this to configs or secrets
	// ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/10eb704e03614a2d989085a1912b6344")
	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/0544b410773d42a38a22c28af4c269b7")
	if err != nil {
		log.Printf("create ethClient failed; err: %v\n", err)
	}
	db, err := ent.Open("sqlite3", "file:/app/backend/data/foo.db?cache=shared&_fk=1")
	if err != nil {
		log.Printf("create db failed; err: %v\n", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	s := &Server{
		db:        db,
		ethClient: ethClient,
		moralisConfig: &moralisConfig{
			url: "https://deep-index.moralis.io/api/v2/",
			key: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6ImEyYzE0MzZlLTQ2ZDMtNDQ1NS05NDRkLTVkMzZjNGQ1ZmU4YiIsIm9yZ0lkIjoiMzMxNTgxIiwidXNlcklkIjoiMzQwOTE0IiwidHlwZUlkIjoiMzMzMWRhNzgtNzU5ZS00MjAxLWI4MmItMWNhZjk4YmMwM2E4IiwidHlwZSI6IlBST0pFQ1QiLCJpYXQiOjE2ODM2MTU1NzgsImV4cCI6NDgzOTM3NTU3OH0.PNIyYhn4sqtzuhEZLcydPkZOQC9yWgevzuawGRyyNMI",
		},
	}
	return s
}

func (s *Server) Shutdown() {
	s.db.Close()
}
