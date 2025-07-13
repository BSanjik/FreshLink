package storage

import (
	"database/sql"
	"services/internal/model/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	query := `
        INSERT INTO users (username, phone, email, password, role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(query, user.Username, user.Phone, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetByPhone(phone string) (*entity.User, error) {
	user := &entity.User{}
	query := `
        SELECT id, username, phone, email, password, role, created_at, updated_at
        FROM users WHERE phone = $1
    `
	err := r.db.QueryRow(query, phone).Scan(&user.ID, &user.Username, &user.Phone, &user.Email,
		&user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *UserRepository) GetByID(id int) (*entity.User, error) {
	user := &entity.User{}
	query := `
        SELECT id, username, phone, email, password, role, created_at, updated_at
        FROM users WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Phone, &user.Email,
		&user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
