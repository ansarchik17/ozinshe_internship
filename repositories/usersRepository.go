package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(connection *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{db: connection}
}


