package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
)

// JDModel holds the schema definition for the JDModel entity.
type JDModel struct {
	ent.Schema
}

// Annotations of the JDModel.
func (JDModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "jd"},
	}
}

// Fields of the JDModel.
func (JDModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("nickname"),
		field.String("password"),
		field.String("email"),
		field.String("id_number"),
		field.Int64("phone_number"),
	}
}

// Edges of the JDModel.
func (JDModel) Edges() []ent.Edge {
	return nil
}
