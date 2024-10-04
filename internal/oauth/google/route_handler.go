package google

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	ent "scaffold/ent/oauth"
	"scaffold/internal/auth"
	"scaffold/internal/oauth"
	"scaffold/pkg/jwt"
	"scaffold/pkg/log"
	oauthHandler "scaffold/pkg/oauth"
)

type RouteHandler struct {
	config *oauth2.Config
	forger *jwt.Forger
	svc    oauth.Service
}

func NewRouteHandler(
	config *oauth2.Config,
	forger *jwt.Forger,
	svc oauth.Service,
) (*RouteHandler, error) {
	if config == nil {
		return nil, errors.New("providing *oauth2.Config is nil")
	}
	if forger == nil {
		return nil, errors.New("providing *forger.Config is nil")
	}
	return &RouteHandler{
		config: config,
		forger: forger,
		svc:    svc,
	}, nil
}

func (h *RouteHandler) AttachOn(router chi.Router) {
	router.Get("/v1/oauth/google", h.Redirect)
	router.Get("/v1/oauth/google/exchange", h.SignIn)
}

func (h *RouteHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	oauthHandler.HandleRedirect(h.config, w, r)
}

func (h *RouteHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	token := oauthHandler.HandleExchange(h.config, w, r)
	if token == nil {
		return
	}

	userInfo, err := h.fetchUserInfo(r.Context(), token)
	if err != nil {
		log.Get(r.Context()).Error("failed to fetch google user info", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Get(r.Context()).Info("google user login", zap.Any("user_info", userInfo))

	user, err := h.svc.LinkAndSignIn(
		r.Context(), ent.IssuerGoogle, userInfo.ID,
		userInfo.Name, userInfo.Email, userInfo.Picture)
	if err != nil {
		log.Get(r.Context()).Error("failed to link and sign in", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userID := strconv.Itoa(user.ID)

	jwtStr, err := h.forger.New(userID)
	if err != nil {
		log.Get(r.Context()).Error("failed to forge jwt", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(auth.LoginResponse{
		Token:       jwtStr,
		UserID:      userID,
		RedirectURL: "/index",
	})
	if err != nil {
		log.Get(r.Context()).Error("failed to encode response", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *RouteHandler) fetchUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error) {
	client := h.config.Client(ctx, token)

	res, err := client.Get(UserInfoURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
