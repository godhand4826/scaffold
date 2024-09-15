package jwt

import (
	"errors"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/rs/xid"
)

var ErrTokenExpired = errors.New("JWT Token has expired")

type Config struct {
	SignKey    string
	ExpireTime time.Duration
}

type Forger struct {
	signOpt    jwt.SignEncryptParseOption
	expireTime time.Duration
}

func NewJWTForger(config Config) *Forger {
	return &Forger{
		signOpt:    jwt.WithKey(jwa.HS256, []byte(config.SignKey)),
		expireTime: config.ExpireTime,
	}
}

func (f *Forger) New(subject string) (string, error) {
	token, err := jwt.
		NewBuilder().
		JwtID(xid.New().String()).
		Subject(subject).
		Expiration(time.Now().Add(f.expireTime)).
		Build()
	if err != nil {
		return "", err
	}

	signed, err := jwt.Sign(token, f.signOpt)
	if err != nil {
		return "", err
	}
	return string(signed), nil
}

func (f *Forger) Verify(token string) (string, error) {
	t, err := jwt.Parse([]byte(token), f.signOpt)
	if err != nil {
		return "", err
	}

	if !t.Expiration().After(time.Now()) {
		return "", ErrTokenExpired
	}

	return t.Subject(), nil
}
