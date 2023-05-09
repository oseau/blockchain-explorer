// Code generated by ent, DO NOT EDIT.

package balance

import (
	"math/big"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

const (
	// Label holds the string label denoting the balance type in the database.
	Label = "balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldBlockNumber holds the string denoting the block_number field in the database.
	FieldBlockNumber = "block_number"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// Table holds the table name of the balance in the database.
	Table = "balances"
)

// Columns holds all SQL columns for balance fields.
var Columns = []string{
	FieldID,
	FieldAccount,
	FieldBlockNumber,
	FieldBalance,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ValueScanner of all Balance fields.
	ValueScanner struct {
		BlockNumber field.TypeValueScanner[*big.Int]
		Balance     field.TypeValueScanner[*big.Int]
	}
)

// OrderOption defines the ordering options for the Balance queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccount orders the results by the account field.
func ByAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccount, opts...).ToFunc()
}

// ByBlockNumber orders the results by the block_number field.
func ByBlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlockNumber, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}
