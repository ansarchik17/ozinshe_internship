package models

import "time"

type MovieDTO struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Director         string    `json:"director"`
	Producer         string    `json:"producer"`
	Year             int32     `json:"year"`
	Timing           int32     `json:"timing"`
	Trend            bool      `json:"trend"`
	Favorite         bool      `json:"favorite"`
	MovieType        string    `json:"movieType"`
	KeyWords         string    `json:"keyWords"`
	WatchCount       int32     `json:"watchCount"`
	SeasonCount      int32     `json:"seasonCount"`
	SeriesCount      int32     `json:"seriesCount"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`

	Categories   []CategoryDTO    `json:"categories"`
	CategoryAges []CategoryAgeDTO `json:"categoryAges"`
	Genres       []GenreDTO       `json:"genres"`
	Poster       PosterDTO        `json:"poster"`
	Screenshots  []ScreenshotDTO  `json:"screenshots"`
	Video        VideoDTO         `json:"video"`
}
