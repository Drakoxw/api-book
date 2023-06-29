package repository

import (
	"database/sql"
	"log"

	"api-book/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FindAllUsers() ([]models.User, error) {
	query := "SELECT id, username, password, email, created_at FROM users"

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Printf("Error al escanear los resultados: %v", err)
			continue
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := ur.db.Exec(query, user.Username, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	query := "UPDATE users SET password = $1, email = $2, updated_at = $3 WHERE username = $4"

	_, err := ur.db.Exec(query, user.Password, user.Email, user.UpdatedAt, user.Username)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(username string) error {
	query := "DELETE FROM users WHERE username = $1"

	_, err := ur.db.Exec(query, username)
	if err != nil {
		return err
	}

	return nil
}
