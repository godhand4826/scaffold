package config

import (
	"golang.org/x/oauth2"

	"scaffold/pkg/logger"
)

type _Config struct {
	Config      string  `mapstructure:"config"`
	FxVerbose   bool    `mapstructure:"fx_verbose"`
	Logger      _Logger `mapstructure:"logger"`
	ServerAddr  string  `mapstructure:"server_addr"`
	GoogleOauth _OAuth  `mapstructure:"google_oauth"`
	GithubOauth _OAuth  `mapstructure:"github_oauth"`
}

func (c _Config) toConfig() *Config {
	return &Config{
		Config:      c.Config,
		FxVerbose:   c.FxVerbose,
		Logger:      c.Logger.toConfig(),
		ServerAddr:  c.ServerAddr,
		GoogleOAuth: c.GoogleOauth.toConfig(),
		GithubOAuth: c.GithubOauth.toConfig(),
	}
}

type _Logger struct {
	Level                string `mapstructure:"level"`
	EnableConsoleEncoder bool   `mapstructure:"enable_console_encoder"`
	EnableCaller         bool   `mapstructure:"enable_caller"`
	EnableUTC            bool   `mapstructure:"enable_utc"`
}

func (c _Logger) toConfig() logger.Config {
	return logger.Config(c)
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
