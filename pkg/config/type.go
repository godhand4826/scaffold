package config

import (
	"time"

	"golang.org/x/oauth2"

	"scaffold/pkg/jwt"
	"scaffold/pkg/log"
	"scaffold/pkg/pg"
)

type _Config struct {
	Config      string    `mapstructure:"config"`
	FxVerbose   bool      `mapstructure:"fx_verbose"`
	Log         _Logger   `mapstructure:"log"`
	ServerAddr  string    `mapstructure:"server_addr"`
	Jwt         _Jwt      `mapstructure:"jwt"`
	GoogleOauth _OAuth    `mapstructure:"google_oauth"`
	GithubOauth _OAuth    `mapstructure:"github_oauth"`
	Postgres    _Postgres `mapstructure:"postgres"`
}

func (c _Config) toConfig() *Config {
	return &Config{
		Config:      c.Config,
		FxVerbose:   c.FxVerbose,
		Log:         c.Log.toConfig(),
		ServerAddr:  c.ServerAddr,
		Jwt:         c.Jwt.toConfig(),
		GoogleOAuth: c.GoogleOauth.toConfig(),
		GithubOAuth: c.GithubOauth.toConfig(),
		Postgres:    c.Postgres.toConfig(),
	}
}

type _Logger struct {
	Level                string `mapstructure:"level"`
	EnableConsoleEncoder bool   `mapstructure:"enable_console_encoder"`
	EnableCaller         bool   `mapstructure:"enable_caller"`
	EnableUTC            bool   `mapstructure:"enable_utc"`
}

func (c _Logger) toConfig() log.Config {
	return log.Config(c)
}

type _Jwt struct {
	SignKey    string        `mapstructure:"sign_key"`
	ExpireTime time.Duration `mapstructure:"expire_time"`
}

func (c _Jwt) toConfig() jwt.Config {
	return jwt.Config(c)
}

type _OAuth struct {
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	RedirectURL  string   `mapstructure:"redirect_url"`
	AuthURL      string   `mapstructure:"auth_url"`
	TokenURL     string   `mapstructure:"token_url"`
	Scopes       []string `mapstructure:"scopes"`
}

func (c _OAuth) toConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthURL,
			TokenURL: c.TokenURL,
		},
		Scopes: c.Scopes,
	}
}

type _Postgres struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

func (c _Postgres) toConfig() pg.Config {
	return pg.Config{
		User:     c.User,
		Password: c.Password,
		Host:     c.Host,
		Port:     c.Port,
		Database: c.Database,
	}
}
