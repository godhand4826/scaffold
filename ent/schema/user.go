package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("avatar"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth", OAuth.Type),
	}
}
