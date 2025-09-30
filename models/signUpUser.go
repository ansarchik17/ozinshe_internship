package models

type SignUpUser struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
}
