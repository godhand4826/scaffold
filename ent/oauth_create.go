// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"scaffold/ent/oauth"
	"scaffold/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OAuthCreate is the builder for creating a OAuth entity.
type OAuthCreate struct {
	config
	mutation *OAuthMutation
	hooks    []Hook
}

// SetIssuer sets the "issuer" field.
func (oc *OAuthCreate) SetIssuer(o oauth.Issuer) *OAuthCreate {
	oc.mutation.SetIssuer(o)
	return oc
}

// SetSubject sets the "subject" field.
func (oc *OAuthCreate) SetSubject(s string) *OAuthCreate {
	oc.mutation.SetSubject(s)
	return oc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (oc *OAuthCreate) SetUserID(id int) *OAuthCreate {
	oc.mutation.SetUserID(id)
	return oc
}

// SetUser sets the "user" edge to the User entity.
func (oc *OAuthCreate) SetUser(u *User) *OAuthCreate {
	return oc.SetUserID(u.ID)
}

// Mutation returns the OAuthMutation object of the builder.
func (oc *OAuthCreate) Mutation() *OAuthMutation {
	return oc.mutation
}

// Save creates the OAuth in the database.
func (oc *OAuthCreate) Save(ctx context.Context) (*OAuth, error) {
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OAuthCreate) SaveX(ctx context.Context) *OAuth {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OAuthCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OAuthCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OAuthCreate) check() error {
	if _, ok := oc.mutation.Issuer(); !ok {
		return &ValidationError{Name: "issuer", err: errors.New(`ent: missing required field "OAuth.issuer"`)}
	}
	if v, ok := oc.mutation.Issuer(); ok {
		if err := oauth.IssuerValidator(v); err != nil {
			return &ValidationError{Name: "issuer", err: fmt.Errorf(`ent: validator failed for field "OAuth.issuer": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "OAuth.subject"`)}
	}
	if len(oc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "OAuth.user"`)}
	}
	return nil
}

func (oc *OAuthCreate) sqlSave(ctx context.Context) (*OAuth, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OAuthCreate) createSpec() (*OAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &OAuth{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(oauth.Table, sqlgraph.NewFieldSpec(oauth.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.Issuer(); ok {
		_spec.SetField(oauth.FieldIssuer, field.TypeEnum, value)
		_node.Issuer = value
	}
	if value, ok := oc.mutation.Subject(); ok {
		_spec.SetField(oauth.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauth.UserTable,
			Columns: []string{oauth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_auth = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OAuthCreateBulk is the builder for creating many OAuth entities in bulk.
type OAuthCreateBulk struct {
	config
	err      error
	builders []*OAuthCreate
}

// Save creates the OAuth entities in the database.
func (ocb *OAuthCreateBulk) Save(ctx context.Context) ([]*OAuth, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*OAuth, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OAuthMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OAuthCreateBulk) SaveX(ctx context.Context) []*OAuth {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OAuthCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
