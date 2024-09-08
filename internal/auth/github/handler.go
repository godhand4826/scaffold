package github

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"

	"scaffold/pkg/oauth"
)

const UserInfoURL = "https://api.github.com/user"

type Handler struct {
	config *oauth2.Config
}

func NewHandler(config *oauth2.Config) (*Handler, error) {
	if config == nil {
		return nil, errors.New("providing *oauth2.Config is nil")
	}
	return &Handler{
		config: config,
	}, nil
}

func (h *Handler) AttachOn(router chi.Router) {
	router.Get("/v1/oauth/github", h.Redirect)
	router.Get("/v1/oauth/github/callback", h.SignIn)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	oauth.HandleRedirect(h.config, w, r)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	token := oauth.HandleExchange(h.config, w, r)
	if token == nil {
		return
	}

	client := h.config.Client(r.Context(), token)

	res, err := client.Get(UserInfoURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// TODO: create account from github_id if not exist
	// TODO: sign in
	_, _ = io.Copy(w, res.Body)
}
