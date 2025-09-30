package models

type UserProfile struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Birthdate   string  `json:"birthdate"`
	Language    string  `json:"language"`
	PhoneNumber string  `json:"phone_number"`
	User        UserMVP `json:"user"`
}
