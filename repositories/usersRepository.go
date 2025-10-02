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

func (r *UsersRepository) Create(c context.Context, user models.SignUpUser) (int, error) {
	var id int
	err := r.db.QueryRow(c, "insert into users(email, password_hash) values ($1, $2) returning id", user.Email, user.PasswordHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UsersRepository) FindByEmail(c context.Context, email string) (models.SignUpUser, error) {
	var user models.SignUpUser
	row := r.db.QueryRow(c, "select id, email, password_hash from users where email = $1", email)
	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash)
	if err != nil {
		return models.SignUpUser{}, err
	}
	return user, err
}

func (r *UsersRepository) FindById(c context.Context, id int) (models.SignUpUser, error) {
	var user models.SignUpUser
	row := r.db.QueryRow(c, "select id, email, password_hash from users where id = $1", id)
	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash)
	if err != nil {
		return models.SignUpUser{}, err
	}
	return user, nil
}

func (r *UsersRepository) CreateProfile(c context.Context, user models.UserProfile) (int, error) {
	var profileId int
	err := r.db.QueryRow(c, "insert into profiles(name, birthdate, language, phone_number, user_id, email) values ($1, $2, $3, $4, $5, $6) returning id",
		user.Name, user.Birthdate, user.Language, user.PhoneNumber, user.User.Id, user.User.Email,
	).Scan(&profileId)
	if err != nil {
		return 0, err
	}
	return profileId, nil
}

func (r *UsersRepository) FindByIdProfile(c context.Context, profileId int) (models.UserProfile, error) {
	var user models.UserProfile
	err := r.db.QueryRow(c, "select id, name, birthdate, language, phone_number, user_id, email from profiles where id = $1", profileId).Scan(&user.Id, &user.Name, &user.Birthdate, &user.Language, &user.PhoneNumber, &user.User.Id, &user.User.Email)
	if err != nil {
		return models.UserProfile{}, err
	}
	return user, nil
}

func (r *UsersRepository) UpdateProfile(c context.Context, id int, updatedProfile models.UserProfile) error {
	tx, err := r.db.Begin(c)
	if err != nil {
		return err
	}
	_, err = tx.Exec(c, "update profiles set name = $1, birthdate = $2, language = $3, phone_number = $4, email = $5 where id = $6",
		updatedProfile.Name, updatedProfile.Birthdate, updatedProfile.Language, updatedProfile.PhoneNumber, updatedProfile.User.Email, id)
	if err != nil {
		return err
	}
	err = tx.Commit(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepository) UpdatePassword(c context.Context, id int, newPasswordHash string) error {
	_, err := r.db.Exec(c,
		"update users set password_hash = $1 where id = $2",
		newPasswordHash, id)
	return err
}
