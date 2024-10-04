package oauth

import (
	"context"

	"scaffold/ent"
	"scaffold/ent/oauth"
)

type Service interface {
	LinkAndSignIn(ctx context.Context, issuer oauth.Issuer, subject string,
		name, email, avatar string) (*ent.User, error)
}

type Repo interface {
	GetUser(ctx context.Context, issuer oauth.Issuer, subject string) (*ent.User, error)
	CreateUser(ctx context.Context, issuer oauth.Issuer, subject string,
		name, email, avatar string) (*ent.User, error)
}

var _ Service = (*service)(nil)

type service struct {
	repo Repo
}

func New(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) LinkAndSignIn(
	ctx context.Context,
	issuer oauth.Issuer, subject string,
	name, email, avatar string,
) (*ent.User, error) {
	user, err := s.repo.GetUser(ctx, issuer, subject)
	if ent.IsNotFound(err) {
		return s.repo.CreateUser(ctx, issuer, subject, name, email, avatar)
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
