package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
)

// QQModel holds the schema definition for the QQModel entity.
type QQModel struct {
	ent.Schema
}

// Annotations of the QQModel.
func (QQModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "qq"},
	}
}

// Fields of the QQModel.
func (QQModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("qq_number"),
		field.Int64("phone_number"),
	}
}

// Edges of the QQModel.
func (QQModel) Edges() []ent.Edge {
	return nil
}
