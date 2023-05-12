package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oseau/blockchain-explorer/ent"
	"github.com/oseau/blockchain-explorer/ent/balance"
)

func (s Server) CreateBalance(account string, blockNumber, balance big.Int) error {
	create := s.db.Balance.Create().
		SetAccount(account).
		SetBlockNumber(&blockNumber).
		SetBalance(&balance)
	_, err := create.Save(context.Background())
	return err
}

func (s Server) GetBalances(account string) ([]*ent.Balance, error) {
	balances, err := s.db.Balance.Query().
		Where(balance.AccountEQ(account)).
		Order(ent.Desc(balance.FieldBlockNumber)).
		Limit(100).
		All(context.Background())
	if err != nil {
		log.Printf("GetBalance failed; err: %v\n", err)
	}
	return balances, nil
}

func (s Server) UpsertBalance(account string, blockNumber, balance *big.Int) error {
	err := s.db.Balance.Create().
		SetAccount(account).
		SetBlockNumber(blockNumber).
		SetBalance(balance).
		OnConflict().
		UpdateBalance().
		Exec(context.Background())
	if err != nil {
		log.Printf("UpsertBalance failed; err: %v\n", err)
	}
	return err
}

func (s Server) GetBalanceRpc(account string) (*big.Int, error) {
	address := common.HexToAddress(account)
	// only current balance available
	// infura errors "project ID does not have access to archive state"
	b, err := s.ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("GetBalanceRpc failed; err: %v\n", err)
	}
	return b, err
}

func (s Server) GetBalanceAtBlockRpc(account string, blockNumber *big.Int) (*big.Int, error) {
	url := fmt.Sprintf("%v%v/balance?chain=eth&to_block=%v", s.moralisConfig.url, account, blockNumber.String())
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", s.moralisConfig.key)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	var resp RespBalanceAt
	if err := json.Unmarshal(body, &resp); err != nil { // Parse []byte to go struct pointer
		log.Printf("GetBalanceAtBlockRpc Can not unmarshal JSON; err: %v\n", err)
		return nil, err
	}
	b := new(big.Int)
	b.SetString(resp.Balance, 10)
	return b, nil
}

func (s Server) GetRecentBalanceChangeBlockNumbersRpc(account string) ([]*big.Int, error) {
	from := time.Now().Add(7 * -24 * time.Hour) // 1 week ago
	url := fmt.Sprintf("%v%v?chain=eth&from_date=%v", s.moralisConfig.url, account, from.Unix())
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", s.moralisConfig.key)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var resp RespRecentBalances
	if err := json.Unmarshal(body, &resp); err != nil { // Parse []byte to go struct pointer
		log.Printf("GetRecentBalanceChangeBlockNumbersRpc Can not unmarshal JSON; err: %v\n", err)
		return nil, err
	}
	blockNumbers := make([]*big.Int, len(resp.Result))
	for i, resp := range resp.Result {
		n := new(big.Int)
		n.SetString(resp.BlockNumber, 10)
		blockNumbers[i] = n
	}
	return blockNumbers, nil
}

type RespBalanceAt struct {
	Balance string `json:"balance"`
}

type RespRecentBalances struct {
	Total    interface{} `json:"total"`
	PageSize int         `json:"page_size"`
	Page     int         `json:"page"`
	Cursor   interface{} `json:"cursor"`
	Result   []struct {
		Hash                     string      `json:"hash"`
		Nonce                    string      `json:"nonce"`
		TransactionIndex         string      `json:"transaction_index"`
		FromAddress              string      `json:"from_address"`
		ToAddress                string      `json:"to_address"`
		Value                    string      `json:"value"`
		Gas                      string      `json:"gas"`
		GasPrice                 string      `json:"gas_price"`
		Input                    string      `json:"input"`
		ReceiptCumulativeGasUsed string      `json:"receipt_cumulative_gas_used"`
		ReceiptGasUsed           string      `json:"receipt_gas_used"`
		ReceiptContractAddress   interface{} `json:"receipt_contract_address"`
		ReceiptRoot              interface{} `json:"receipt_root"`
		ReceiptStatus            string      `json:"receipt_status"`
		BlockTimestamp           time.Time   `json:"block_timestamp"`
		BlockNumber              string      `json:"block_number"`
		BlockHash                string      `json:"block_hash"`
		TransferIndex            []int       `json:"transfer_index"`
	} `json:"result"`
}
