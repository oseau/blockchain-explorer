package data

import (
	"math"
	"math/big"
	"testing"

	"github.com/oseau/blockchain-explorer/ent"
	"github.com/oseau/blockchain-explorer/ent/enttest"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetBalance(t *testing.T) {
	c := NewClient()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	c.db = client
	if _, err := c.GetBalance("account not existing"); !ent.IsNotFound(err) {
		t.Fatalf("get balance failed; err: %v", err)
	}
	if err := c.CreateBalance("test account", *big.NewInt(1), *big.NewInt(2)); err != nil {
		t.Fatalf("create balance failed; err: %v", err)
	}
	balance, err := c.GetBalance("test account")
	if err != nil {
		t.Fatalf("get balance failed; err: %v", err)
	}
	if balance.Account != "test account" {
		t.Fatalf("unexpected balance account: %v", balance.Account)
	}
}

func TestGetBalanceRpc(t *testing.T) {
	c := NewClient()
	b, err := c.GetBalanceRpc("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		t.Fatalf("get balance rpc failed; err: %v", err)
	}
	balance := new(big.Float)
	balance.SetString(b.String())
	eth := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))
	if eth.Cmp(big.NewFloat(0)) == 0 {
		t.Fatal("0 eth in vitalik.eth")
	}
}
