package handlers

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func NewSignInHandler() *SignInResponse {
	return &SignInResponse{}
}
