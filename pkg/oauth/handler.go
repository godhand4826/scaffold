package oauth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"scaffold/pkg/log"
	"time"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const (
	StateCookieName = "oauth_state"
	StateQueryKey   = "state"
	CodeQueryKey    = "code"
)

func HandleRedirect(config *oauth2.Config, w http.ResponseWriter, r *http.Request) {
	state, err := randomState()
	if err != nil {
		log.Get(r.Context()).Error("failed to generate oauth random state", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	url := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	http.SetCookie(w, &http.Cookie{
		Name:     StateCookieName,
		Value:    state,
		Expires:  time.Now().Add(time.Minute),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, url, http.StatusFound)
}

func HandleExchange(config *oauth2.Config, w http.ResponseWriter, r *http.Request) *oauth2.Token {
	state := r.URL.Query().Get(StateQueryKey)
	cookie, err := r.Cookie(StateCookieName)
	if err != nil || cookie.Value != state {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return nil
	}

	code := r.URL.Query().Get(CodeQueryKey)
	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		log.Get(r.Context()).Error("failed to exchange access token", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	}

	return token
}

func randomState() (string, error) {
	b := make([]byte, 30)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
