package models

type VideoDTO struct {
	Id       int64  `json:"id"`
	Link     string `json:"link"`
	Number   int32  `json:"number"`
	SeasonId int64  `json:"seasonId"`
}
