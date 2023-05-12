package web

import "github.com/oseau/blockchain-explorer/ent"

func toBalance(b *ent.Balance) Balance {
	return Balance{BlockNumber: b.BlockNumber.String(), Balance: b.Balance.String()}
}

func toBalances(bs []*ent.Balance) []Balance {
	result := make([]Balance, len(bs))
	for key, value := range bs {
		result[key] = toBalance(value)
	}
	return result
}
