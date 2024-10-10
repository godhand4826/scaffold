package auth

type LoginResponse struct {
	Token       string `json:"token,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
}
