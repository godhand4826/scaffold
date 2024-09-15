package auth

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	UserID      string `json:"user_id"`
}
