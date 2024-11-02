package repository

import (
	"database/sql"
	"github.com/bohexists/book-crud-svc/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Role,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
