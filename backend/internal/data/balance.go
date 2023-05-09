package data

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oseau/blockchain-explorer/ent"
	"github.com/oseau/blockchain-explorer/ent/balance"
)

func (c Client) CreateBalance(account string, blockNumber, balance big.Int) error {
	create := c.db.Balance.Create().
		SetAccount(account).
		SetBlockNumber(&blockNumber).
		SetBalance(&balance)
	_, err := create.Save(context.Background())
	return err
}

func (c Client) GetBalance(account string) (*ent.Balance, error) {
	return c.db.Balance.Query().Where(balance.AccountEQ(account)).Only(context.Background())
}

func (c Client) GetBalanceRpc(account string) (*big.Int, error) {
	address := common.HexToAddress(account)
	return c.ethClient.BalanceAt(context.Background(), address, nil)
}
