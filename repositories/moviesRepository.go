package repositories

import (
	"context"
	"ozinshe/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MoviesRepository struct {
	db *pgxpool.Pool
}

func NewMovieRepository(connection *pgxpool.Pool) *MoviesRepository {
	return &MoviesRepository{db: connection}
}

func (r *MoviesRepository) Create(c context.Context, movie models.MovieDTO) (int, error) {
	var id int
	err := r.db.QueryRow(c, "insert into movies(name, description, director, producer, year, timing, trend, favorite, movieType, keyWords, watchCount, seasonCount, seriesCount, createdDate, lastModifiedDate) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) returning id", movie.Name, movie.CreatedDate, movie.SeasonCount, movie.LastModifiedDate, movie.WatchCount, movie.KeyWords, movie.MovieType, movie.Timing, movie.Description, movie.Director, movie.Producer, movie.Year, movie.Trend, movie.Favorite, movie.MovieType).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
