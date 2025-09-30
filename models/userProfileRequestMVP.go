package models

type СreateUserProfileRequest struct {
	Id             int    `json:"id"`
	PrivateInfo    string `json:"private_info"`
	ChangePassword string `json:"change_password"`
}
