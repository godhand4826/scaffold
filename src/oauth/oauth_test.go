package oauth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"scaffold/ent"
	oa "scaffold/ent/oauth"
	mockoauth "scaffold/mocks/src/oauth"
	"scaffold/src/oauth"
)

func TestOAuthSuite(t *testing.T) {
	suite.Run(t, new(OAuthSuite))
}

type OAuthSuite struct {
	suite.Suite

	// mocks
	repo *mockoauth.MockRepo

	// SUT
	svc oauth.Service
}

func (s *OAuthSuite) SetupTest() {
	// initialize all mocks
	s.repo = mockoauth.NewMockRepo(s.T())

	// initialize SUT
	s.svc = oauth.New(s.repo)
}

func (s *OAuthSuite) TearDownTest() {
	s.repo = nil
	s.svc = nil
}

type MockExpect struct {
	Repo *RepoMockExpect
}

func (s *OAuthSuite) setupMockExpect(expect MockExpect) {
	expect.Repo.applyOn(s.repo)
}

type RepoMockExpect struct {
	GetUser    *RepoGetUserMockExpect
	CreateUser *RepoCreateUserMockExpect
}

func (me *RepoMockExpect) applyOn(repo *mockoauth.MockRepo) {
	if me == nil {
		return
	}

	me.GetUser.applyOn(repo)
	me.CreateUser.applyOn(repo)
}

type RepoGetUserMockExpect struct {
	Issuer  oa.Issuer
	Subject string

	ReturnUser *ent.User
	ReturnErr  error
}

func (me *RepoGetUserMockExpect) applyOn(repo *mockoauth.MockRepo) {
	if me == nil {
		return
	}

	repo.EXPECT().
		GetUser(mock.Anything, me.Issuer, me.Subject).
		Return(me.ReturnUser, me.ReturnErr).
		Once()
}

type RepoCreateUserMockExpect struct {
	Issuer  oa.Issuer
	Subject string
	Name    string
	Email   string
	Avatar  string

	ReturnUser *ent.User
	ReturnErr  error
}

func (me *RepoCreateUserMockExpect) applyOn(repo *mockoauth.MockRepo) {
	if me == nil {
		return
	}

	repo.EXPECT().
		CreateUser(mock.Anything, me.Issuer, me.Subject, me.Name, me.Email, me.Avatar).
		Return(me.ReturnUser, me.ReturnErr).
		Once()
}

func (s *OAuthSuite) TestLinkAndSignIn() {
	type Arg struct {
		Issuer  oa.Issuer
		Subject string
		Name    string
		Email   string
		Avatar  string
	}
	type Expect struct {
		User  *ent.User
		Error error
	}
	type Test struct {
		Arg        Arg
		MockExpect MockExpect
		Expect     Expect
	}

	const name = "Eric"
	const subject = "subject_123"

	var arg = Arg{
		Issuer:  oa.IssuerGoogle,
		Subject: subject,
		Name:    name,
	}
	var user = ent.User{
		ID:   1,
		Name: arg.Name,
	}
	var databaseErr = errors.New("database error")

	var tests = []Test{
		{
			Arg: arg,
			MockExpect: MockExpect{
				Repo: &RepoMockExpect{
					GetUser: &RepoGetUserMockExpect{
						Issuer:     oa.IssuerGoogle,
						Subject:    subject,
						ReturnUser: &user,
					},
				},
			},
			Expect: Expect{
				User: &user,
			},
		},
		{
			Arg: arg,
			MockExpect: MockExpect{
				Repo: &RepoMockExpect{
					GetUser: &RepoGetUserMockExpect{
						Issuer:    oa.IssuerGoogle,
						Subject:   subject,
						ReturnErr: databaseErr,
					},
				},
			},
			Expect: Expect{
				Error: databaseErr,
			},
		},
		{
			Arg: arg,
			MockExpect: MockExpect{
				Repo: &RepoMockExpect{
					GetUser: &RepoGetUserMockExpect{
						Issuer:    oa.IssuerGoogle,
						Subject:   subject,
						ReturnErr: &ent.NotFoundError{},
					},
					CreateUser: &RepoCreateUserMockExpect{
						Issuer:     oa.IssuerGoogle,
						Subject:    subject,
						Name:       name,
						ReturnUser: &user,
					},
				},
			},
			Expect: Expect{
				User: &user,
			},
		},
		{
			Arg: arg,
			MockExpect: MockExpect{
				Repo: &RepoMockExpect{
					GetUser: &RepoGetUserMockExpect{
						Issuer:    oa.IssuerGoogle,
						Subject:   subject,
						ReturnErr: &ent.NotFoundError{},
					},
					CreateUser: &RepoCreateUserMockExpect{
						Issuer:    oa.IssuerGoogle,
						Subject:   subject,
						Name:      name,
						ReturnErr: databaseErr,
					},
				},
			},
			Expect: Expect{
				Error: databaseErr,
			},
		},
	}
	for _, t := range tests {
		s.setupMockExpect(t.MockExpect)

		u, err := s.svc.LinkAndSignIn(
			context.Background(),
			t.Arg.Issuer,
			t.Arg.Subject,
			t.Arg.Name,
			t.Arg.Email,
			t.Arg.Avatar,
		)

		s.Assert().Equal(t.Expect.User, u)
		s.Assert().Equal(t.Expect.Error, err)
	}
}
