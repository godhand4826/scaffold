package repo

import (
	"context"

	"scaffold/ent"
	"scaffold/ent/oauth"
	"scaffold/pkg/pg"
	oa "scaffold/src/oauth"
)

var _ oa.Repo = (*Repo)(nil)

type Repo struct {
	client *ent.Client
}

func New(client *ent.Client) *Repo {
	return &Repo{
		client: client,
	}
}

func (r *Repo) CreateUser(
	ctx context.Context,
	issuer oauth.Issuer, subject string,
	name string, email string, avatar string,
) (user *ent.User, err error) {
	if err = pg.WithTx(ctx, r.client, func(tx *ent.Tx) error {
		user, err = tx.User.
			Create().
			SetName(name).
			SetEmail(email).
			SetAvatar(avatar).
			Save(ctx)
		if err != nil {
			return err
		}

		if _, err = tx.OAuth.
			Create().
			SetIssuer(issuer).
			SetSubject(subject).
			SetUser(user).
			Save(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetUser(ctx context.Context, issuer oauth.Issuer, subject string) (*ent.User, error) {
	user, err := r.client.OAuth.
		Query().
		Where(oauth.IssuerEQ(issuer), oauth.Subject(subject)).
		QueryUser().
		Only(ctx)
	return user, err
}
