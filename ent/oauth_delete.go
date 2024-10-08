// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"scaffold/ent/oauth"
	"scaffold/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OAuthDelete is the builder for deleting a OAuth entity.
type OAuthDelete struct {
	config
	hooks    []Hook
	mutation *OAuthMutation
}

// Where appends a list predicates to the OAuthDelete builder.
func (od *OAuthDelete) Where(ps ...predicate.OAuth) *OAuthDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OAuthDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OAuthDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OAuthDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(oauth.Table, sqlgraph.NewFieldSpec(oauth.FieldID, field.TypeInt))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OAuthDeleteOne is the builder for deleting a single OAuth entity.
type OAuthDeleteOne struct {
	od *OAuthDelete
}

// Where appends a list predicates to the OAuthDelete builder.
func (odo *OAuthDeleteOne) Where(ps ...predicate.OAuth) *OAuthDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OAuthDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oauth.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OAuthDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}
