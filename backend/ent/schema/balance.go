package schema

import (
	"math/big"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Balance holds the schema definition for the Balance entity.
type Balance struct {
	ent.Schema
}

// Fields of the Balance.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		field.String("account"),
		field.String("block_number").
			GoType(&big.Int{}).ValueScanner(field.TextValueScanner[*big.Int]{}),
		field.String("balance").
			GoType(&big.Int{}).ValueScanner(field.TextValueScanner[*big.Int]{}),
	}
}

func (Balance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account"),
		index.Fields("account", "block_number").Unique(),
	}
}

// Edges of the Balance.
func (Balance) Edges() []ent.Edge {
	return nil
}
