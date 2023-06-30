package repository

import (
	"database/sql"
	"errors"
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

func (ur *UserRepository) GetUserId(id int) (models.User, error) {
	query := "SELECT id, username, password, email, created_at FROM users WHERE id = $1"

	row := ur.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("no found User")
		}
		return user, err
	}

	return user, nil
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
	query := "INSERT INTO users (username, password, email, created_at) VALUES ($1, $2, $3, $4)"

	_, err := ur.db.Exec(query, user.Username, user.Password, user.Email, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUser(user *models.UpdateUser, id int) error {
	query := "UPDATE users SET email = $1, updated_at = $2 WHERE id = $3"

	_, err := ur.db.Exec(query, user.Email, user.UpdatedAt, id)
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
