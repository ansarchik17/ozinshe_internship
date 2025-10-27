package repositories

import (
	"context"
	"ozinshe/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(connection *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: connection}
}

func (r *CategoryRepository) Create(c context.Context, category models.CategoryDTO) (int, error) {
	var id int
	err := r.db.QueryRow(c, "insert into categories(name, id, link, fileId, movieCount) values ($1, $2, $3, $4, $5) returning id",
		category.Name,
		category.Id,
		category.Link,
		category.FileId,
		category.MovieCount,
	).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
