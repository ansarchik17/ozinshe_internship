package models

type PosterDTO struct {
	Id      int64  `json:"id"`
	FileId  int64  `json:"fileId"`
	Link    string `json:"link"`
	MovieId int64  `json:"movieId"`
}
