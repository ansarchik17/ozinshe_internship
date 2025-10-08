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
	err := r.db.QueryRow(c,
		"INSERT INTO movies(name, description, director, producer, year, timing, trend, favorite, movie_type, key_words, watch_count, season_count, series_count, created_date, last_modified_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id",
		movie.Name,
		movie.Description,
		movie.Director,
		movie.Producer,
		movie.Year,
		movie.Timing,
		movie.Trend,
		movie.Favorite,
		movie.MovieType,
		movie.KeyWords,
		movie.WatchCount,
		movie.SeasonCount,
		movie.SeriesCount,
		movie.CreatedDate,
		movie.LastModifiedDate,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
