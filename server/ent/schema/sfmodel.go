package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
)

// SFModel holds the schema definition for the SFModel entity.
type SFModel struct {
	ent.Schema
}

// Annotations of the SFModel.
func (SFModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sf"},
	}
}

// Fields of the SFModel.
func (SFModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("phone_number"),
		field.String("address"),
	}
}

// Edges of the SFModel.
func (SFModel) Edges() []ent.Edge {
	return nil
}
