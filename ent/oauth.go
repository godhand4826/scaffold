// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"scaffold/ent/oauth"
	"scaffold/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OAuth is the model entity for the OAuth schema.
type OAuth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Issuer holds the value of the "issuer" field.
	Issuer oauth.Issuer `json:"issuer,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OAuthQuery when eager-loading is set.
	Edges        OAuthEdges `json:"edges"`
	user_auth    *int
	selectValues sql.SelectValues
}

// OAuthEdges holds the relations/edges for other nodes in the graph.
type OAuthEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OAuthEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuth) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauth.FieldID:
			values[i] = new(sql.NullInt64)
		case oauth.FieldIssuer, oauth.FieldSubject:
			values[i] = new(sql.NullString)
		case oauth.ForeignKeys[0]: // user_auth
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuth fields.
func (o *OAuth) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case oauth.FieldIssuer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field issuer", values[i])
			} else if value.Valid {
				o.Issuer = oauth.Issuer(value.String)
			}
		case oauth.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				o.Subject = value.String
			}
		case oauth.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_auth", value)
			} else if value.Valid {
				o.user_auth = new(int)
				*o.user_auth = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OAuth.
// This includes values selected through modifiers, order, etc.
func (o *OAuth) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the OAuth entity.
func (o *OAuth) QueryUser() *UserQuery {
	return NewOAuthClient(o.config).QueryUser(o)
}

// Update returns a builder for updating this OAuth.
// Note that you need to call OAuth.Unwrap() before calling this method if this OAuth
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *OAuth) Update() *OAuthUpdateOne {
	return NewOAuthClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the OAuth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *OAuth) Unwrap() *OAuth {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuth is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *OAuth) String() string {
	var builder strings.Builder
	builder.WriteString("OAuth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("issuer=")
	builder.WriteString(fmt.Sprintf("%v", o.Issuer))
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(o.Subject)
	builder.WriteByte(')')
	return builder.String()
}

// OAuths is a parsable slice of OAuth.
type OAuths []*OAuth
