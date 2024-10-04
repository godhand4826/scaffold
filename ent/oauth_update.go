// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"scaffold/ent/oauth"
	"scaffold/ent/predicate"
	"scaffold/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OAuthUpdate is the builder for updating OAuth entities.
type OAuthUpdate struct {
	config
	hooks    []Hook
	mutation *OAuthMutation
}

// Where appends a list predicates to the OAuthUpdate builder.
func (ou *OAuthUpdate) Where(ps ...predicate.OAuth) *OAuthUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *OAuthUpdate) SetUserID(id int) *OAuthUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *OAuthUpdate) SetUser(u *User) *OAuthUpdate {
	return ou.SetUserID(u.ID)
}

// Mutation returns the OAuthMutation object of the builder.
func (ou *OAuthUpdate) Mutation() *OAuthMutation {
	return ou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ou *OAuthUpdate) ClearUser() *OAuthUpdate {
	ou.mutation.ClearUser()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OAuthUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OAuthUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OAuthUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OAuthUpdate) check() error {
	if ou.mutation.UserCleared() && len(ou.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OAuth.user"`)
	}
	return nil
}

func (ou *OAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauth.Table, oauth.Columns, sqlgraph.NewFieldSpec(oauth.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ou.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OAuthUpdateOne is the builder for updating a single OAuth entity.
type OAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OAuthMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *OAuthUpdateOne) SetUserID(id int) *OAuthUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *OAuthUpdateOne) SetUser(u *User) *OAuthUpdateOne {
	return ouo.SetUserID(u.ID)
}

// Mutation returns the OAuthMutation object of the builder.
func (ouo *OAuthUpdateOne) Mutation() *OAuthMutation {
	return ouo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *OAuthUpdateOne) ClearUser() *OAuthUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// Where appends a list predicates to the OAuthUpdate builder.
func (ouo *OAuthUpdateOne) Where(ps ...predicate.OAuth) *OAuthUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OAuthUpdateOne) Select(field string, fields ...string) *OAuthUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated OAuth entity.
func (ouo *OAuthUpdateOne) Save(ctx context.Context) (*OAuth, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OAuthUpdateOne) SaveX(ctx context.Context) *OAuth {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OAuthUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OAuthUpdateOne) check() error {
	if ouo.mutation.UserCleared() && len(ouo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OAuth.user"`)
	}
	return nil
}

func (ouo *OAuthUpdateOne) sqlSave(ctx context.Context) (_node *OAuth, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauth.Table, oauth.Columns, sqlgraph.NewFieldSpec(oauth.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth.FieldID)
		for _, f := range fields {
			if !oauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ouo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OAuth{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
