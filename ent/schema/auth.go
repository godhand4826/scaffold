package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OAuth struct {
	ent.Schema
}

func (OAuth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "oauth"},
	}
}

func (OAuth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("issuer", "subject").
			Unique(),
	}
}

func (OAuth) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("issuer").Values("google", "github").Immutable(),
		field.String("subject").Immutable(),
	}
}

func (OAuth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auth").
			Required().
			Unique(),
	}
}
