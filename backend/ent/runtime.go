// Code generated by ent, DO NOT EDIT.

package ent

import (
	"math/big"

	"entgo.io/ent/schema/field"
	"github.com/oseau/blockchain-explorer/ent/balance"
	"github.com/oseau/blockchain-explorer/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	balanceFields := schema.Balance{}.Fields()
	_ = balanceFields
	// balanceDescBlockNumber is the schema descriptor for block_number field.
	balanceDescBlockNumber := balanceFields[1].Descriptor()
	balance.ValueScanner.BlockNumber = balanceDescBlockNumber.ValueScanner.(field.TypeValueScanner[*big.Int])
	// balanceDescBalance is the schema descriptor for balance field.
	balanceDescBalance := balanceFields[2].Descriptor()
	balance.ValueScanner.Balance = balanceDescBalance.ValueScanner.(field.TypeValueScanner[*big.Int])
}
