package models

type CategoryAgeDTO struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	FileId     int64  `json:"fileId"`
	MovieCount int32  `json:"movieCount"`
}
