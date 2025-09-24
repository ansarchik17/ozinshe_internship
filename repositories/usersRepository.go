package repositories

import (
	"context"
	"ozinshe/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(conn *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{db: conn}
}

func (r *UsersRepository) Create(c context.Context, user models.User) (int, error) {
	var id int
	err := r.db.QueryRow(c, "insert into users(email, password_hash, avatar_url) values ($1, $2, $3) returning id", user.Email, user.PasswordHash, user.AvatarUrl).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UsersRepository) FindByEmail(c context.Context, email string) (models.User, error) {
	var user models.User
	row := r.db.QueryRow(c, "select id, email, password_hash, avatar_url from users where email = $1", email)
	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.AvatarUrl)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}
