package handlers

import (
	"net/http"
	"ozinshe/models"
	"ozinshe/repositories"

	"github.com/gin-gonic/gin"
)

type MoviesHandler struct {
	movieRepo *repositories.MoviesRepository
}

func NewMovieHandler(movieRepo *repositories.MoviesRepository) *MoviesHandler {
	return &MoviesHandler{movieRepo: movieRepo}
}

func (handler *MoviesHandler) CreateMovie(c *gin.Context) {
	var request models.MovieDTO
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could bind json object"))
		return
	}
	movie := models.MovieDTO{
		Name:             request.Name,
		Description:      request.Description,
		Director:         request.Director,
		Producer:         request.Producer,
		Year:             request.Year,
		Timing:           request.Timing,
		Trend:            request.Trend,
		Favorite:         request.Favorite,
		MovieType:        request.MovieType,
		KeyWords:         request.KeyWords,
		WatchCount:       request.WatchCount,
		SeasonCount:      request.SeasonCount,
		SeriesCount:      request.SeriesCount,
		CreatedDate:      request.CreatedDate,
		LastModifiedDate: request.LastModifiedDate,
	}
	id, err := handler.movieRepo.Create(c, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewApiError("Error creating movie"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

//Ansar
