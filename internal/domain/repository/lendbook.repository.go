package repository

import (
	"api-book/internal/domain/models"
	"database/sql"
	"fmt"
	"time"
)

type LendBookRepository struct {
	db *sql.DB
}

func NewLendBookRepository(db *sql.DB) *LendBookRepository {
	return &LendBookRepository{
		db: db,
	}
}

/** handler.CreateLendBook */
func (lbr *LendBookRepository) CreateLendBook(lendBook *models.LendBook) error {
	query := "INSERT INTO lend_books (user_id, book_id, created_at) VALUES ($1, $2, $3)"

	_, err := lbr.db.Exec(query, lendBook.UserId, lendBook.BookId, time.Now())
	if err != nil {
		return fmt.Errorf("error al crear el préstamo de libro: %v", err)
	}

	return nil
}

/** handler.GetLendBook */
func (lbr *LendBookRepository) GetLendBookByID(id int) (*models.LendBook, error) {
	query := "SELECT id, user_id, book_id, return_book, created_at, updated_at FROM lend_books WHERE id = $1"

	row := lbr.db.QueryRow(query, id)

	var lendBook models.LendBook
	err := row.Scan(&lendBook.Id, &lendBook.UserId, &lendBook.BookId, &lendBook.ReturnBook, &lendBook.CreatedAt, &lendBook.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("préstamo de libro no encontrado")
		}
		return nil, err
	}

	return &lendBook, nil
}

/** handler. */
func (lbr *LendBookRepository) ReturnBookToLibrary(bookId int) error {
	query := "UPDATE lend_books SET return_book = $1, updated_at = $2 WHERE id = $3"

	_, err := lbr.db.Exec(query, time.Now(), time.Now(), bookId)
	if err != nil {
		return fmt.Errorf("error al actualizar la fecha de retorno del libro: %v", err)
	}

	return nil
}

/** handler.ListLendBooks */
func (lbr *LendBookRepository) GetAllBooksAndLends() ([]models.Book, error) {
	query := `
		SELECT b.id, b.title, b.author, b.literary_genre, lb.id, lb.user_id, lb.return_book, lb.created_at, lb.updated_at
		FROM books b
		LEFT JOIN lend_books lb ON b.id = lb.book_id
		ORDER BY b.id ASC, lb.created_at ASC
	`

	rows, err := lbr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los libros y prestamos: %v", err)
	}
	defer rows.Close()

	var books []models.Book
	currentBook := models.Book{}
	var lendBook *models.LendBook

	for rows.Next() {
		var (
			bookID         int64
			bookTitle      string
			bookAuthor     string
			bookGenre      string
			lendBookID     sql.NullInt64
			userID         sql.NullInt64
			returnBookTime sql.NullTime
			createdAt      time.Time
			updatedAt      time.Time
		)

		err := rows.Scan(
			&bookID,
			&bookTitle,
			&bookAuthor,
			&bookGenre,
			&lendBookID,
			&userID,
			&returnBookTime,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los resultados: %v", err)
		}

		if lendBookID.Valid {
			lendBook = &models.LendBook{
				Id:         lendBookID.Int64,
				UserId:     userID.Int64,
				BookId:     bookID,
				ReturnBook: returnBookTime,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
			}
		} else {
			lendBook = nil
		}

		if currentBook.Id != bookID {
			currentBook = models.Book{
				Id:            bookID,
				Title:         bookTitle,
				Author:        bookAuthor,
				LiteraryGenre: bookGenre,
				LendHistory:   make([]*models.LendBook, 0),
			}
			books = append(books, currentBook)
		}

		if lendBook != nil {
			currentBook.LendHistory = append(currentBook.LendHistory, lendBook)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar por los resultados: %v", err)
	}

	return books, nil
}

/** handler. */
func (lbr *LendBookRepository) GetAllUsersAndLends() ([]models.User, error) {
	query := `
		SELECT u.id, u.username, u.email, lb.id, lb.book_id, lb.return_book, lb.created_at, lb.updated_at
		FROM users u
		LEFT JOIN lend_books lb ON u.id = lb.user_id
		ORDER BY u.id ASC, lb.created_at ASC
	`

	rows, err := lbr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los usuarios y préstamos: %v", err)
	}
	defer rows.Close()

	var users []models.User
	currentUser := models.User{}
	var lendBook *models.LendBook

	for rows.Next() {
		var (
			userID         int64
			username       string
			email          string
			lendBookID     sql.NullInt64
			bookID         sql.NullInt64
			returnBookTime sql.NullTime
			createdAt      time.Time
			updatedAt      time.Time
		)

		err := rows.Scan(
			&userID,
			&username,
			&email,
			&lendBookID,
			&bookID,
			&returnBookTime,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los resultados: %v", err)
		}

		if lendBookID.Valid {
			lendBook = &models.LendBook{
				Id:         lendBookID.Int64,
				UserId:     userID,
				BookId:     bookID.Int64,
				ReturnBook: returnBookTime,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
			}
		} else {
			lendBook = nil
		}

		if currentUser.Id != userID {
			currentUser = models.User{
				Id:          userID,
				Username:    username,
				Email:       email,
				LendHistory: make([]*models.LendBook, 0),
			}
			users = append(users, currentUser)
		}

		if lendBook != nil {
			currentUser.LendHistory = append(currentUser.LendHistory, lendBook)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar por los resultados: %v", err)
	}

	return users, nil
}
