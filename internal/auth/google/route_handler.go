package google

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"

	"scaffold/pkg/oauth"
)

const UserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"

type RouteHandler struct {
	config *oauth2.Config
}

func NewRouteHandler(config *oauth2.Config) (*RouteHandler, error) {
	if config == nil {
		return nil, errors.New("providing *oauth2.Config is nil")
	}
	return &RouteHandler{
		config: config,
	}, nil
}

func (h *RouteHandler) AttachOn(router chi.Router) {
	router.Get("/v1/oauth/google", h.Redirect)
	router.Get("/v1/oauth/google/callback", h.SignIn)
}

func (h *RouteHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	oauth.HandleRedirect(h.config, w, r)
}

func (h *RouteHandler) SignIn(w http.ResponseWriter, r *http.Request) {
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

	// TODO: create account from google_id if not exist
	// TODO: sign in
	_, _ = io.Copy(w, res.Body)
}
