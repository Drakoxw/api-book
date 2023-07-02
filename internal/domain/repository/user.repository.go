package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

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

func (ur *UserRepository) GetUsersHistory() ([]*models.UserHistory, error) {
	query := `
		SELECT 
			u.id, 
			u.username, 
			u.email,
			l.id AS lend_id, 
			l.user_id, 
			l.book_id, 
			l.return_book, 
			l.created_at AS lend_created_at, 
			l.updated_at
		FROM users AS u
		JOIN lend_books AS l ON u.id = l.user_id
		ORDER BY u.id DESC
	`
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersMap := make(map[int64]*models.UserHistory)

	for rows.Next() {
		var id, lendId, bookId, userId int64
		var username, email string
		var returnBook sql.NullTime
		var lendCreated, lendUpdated time.Time

		err := rows.Scan(
			&id,
			&username,
			&email,
			&lendId,
			&userId,
			&bookId,
			&returnBook,
			&lendCreated,
			&lendUpdated,
		)
		if err != nil {
			return nil, err
		}

		user, ok := usersMap[id]
		if !ok {
			user = &models.UserHistory{
				Id:          id,
				Email:       email,
				Username:    username,
				LendHistory: []*models.LendBook{},
			}
			usersMap[id] = user
		}

		lend := &models.LendBook{
			Id:         lendId,
			UserId:     userId,
			BookId:     bookId,
			ReturnBook: returnBook,
			CreatedAt:  lendCreated,
			UpdatedAt:  lendUpdated,
		}
		user.LendHistory = append(user.LendHistory, lend)

	}
	users := make([]*models.UserHistory, 0, len(usersMap))
	for _, user := range usersMap {
		users = append(users, user)
	}

	return users, nil
}
