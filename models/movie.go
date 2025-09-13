package models

type Movie struct {
	Id          int
	Title       string
	Description string
	PosterUrl   string
	Genre       string
	ReleaseDate string
	Rating      int
	Duration    int
}
