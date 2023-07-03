package repository

import (
	"api-book/internal/domain/models"
	"api-book/internal/infrastructure/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) CreateBook(book *models.Book) error {
	query := "INSERT INTO books (title, author, literary_genre, created_at) VALUES ($1, $2, $3, $4)"

	_, err := br.db.Exec(query, book.Title, book.Author, book.LiteraryGenre, time.Now())
	if err != nil {
		return fmt.Errorf("error al crear el libro: %v", err)
	}

	return nil
}

func (br *BookRepository) GetBookByID(id int) (*models.Book, error) {
	query := "SELECT id, title, author, literary_genre, created_at FROM books WHERE id = $1"

	row := br.db.QueryRow(query, id)

	var book models.Book
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.LiteraryGenre, &book.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("libro no encontrado")
		}
		return nil, err
	}

	return &book, nil
}

func (br *BookRepository) ListBooks(page int, limit int) ([]models.Book, error) {
	query := `
		SELECT id, title, author, literary_genre, created_at
		FROM books 
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	rows, err := br.db.Query(query, limit, page)
	if err != nil {
		return nil, fmt.Errorf("error al obtener la lista de libros: %v", err)
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.LiteraryGenre,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los resultados: %v", err)
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar por los resultados: %v", err)
	}

	return books, nil
}

func (br *BookRepository) GetBooksHistory() ([]*models.Book, error) {

	query := `
		SELECT 
			b.id, 
			b.title, 
			b.author, 
			b.literary_genre, 
			b.created_at,
			l.id AS lend_id, 
			l.user_id, 
			l.book_id, 
			l.return_book, 
			l.created_at AS lend_created_at, 
			l.updated_at
		FROM books AS b
		JOIN lend_books AS l ON b.id = l.book_id
		ORDER BY b.id DESC
	`

	rows, err := br.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookMap := make(map[int64]*models.Book)

	for rows.Next() {
		var bookID, lendID sql.NullInt64
		var title, author, literaryGenre string
		var createdAt time.Time
		var lendCreatedAt, updatedAt, returnBook sql.NullTime
		var userID, bookLendId int64

		err := rows.Scan(
			&bookID,
			&title,
			&author,
			&literaryGenre,
			&createdAt,
			&lendID,
			&userID,
			&bookLendId,
			&returnBook,
			&lendCreatedAt,
			&updatedAt)
		if err != nil {
			return nil, err
		}

		// Verificar si el libro ya existe en el mapa, de lo contrario, crearlo
		book, ok := bookMap[bookLendId]
		if !ok {
			book = &models.Book{
				Id:            bookLendId,
				Title:         title,
				Author:        author,
				LiteraryGenre: literaryGenre,
				CreatedAt:     createdAt,
				LendHistory:   []*models.LendBook{},
			}
			bookMap[bookLendId] = book
		}

		// Si hay información de préstamo, agregarla al historial de préstamos del libro
		if lendID.Valid {
			lend := &models.LendBook{
				Id:         lendID.Int64,
				UserId:     userID,
				BookId:     bookLendId,
				ReturnBook: returnBook,
				CreatedAt:  lendCreatedAt.Time,
				UpdatedAt:  updatedAt.Time,
			}
			book.LendHistory = append(book.LendHistory, lend)
		}
	}

	books := make([]*models.Book, 0, len(bookMap))
	for _, book := range bookMap {
		books = append(books, book)
	}

	return books, nil
}

func (br *BookRepository) GetBooksHistoryV2() ([]models.BookV2, error) {

	query := `
	SELECT 
    b.id, 
    b.title, 
    b.author, 
    b.literary_genre, 
    b.created_at,
		(
			SELECT json_agg(json_build_object(
				'id', l.id,
				'user_id', l.user_id,
				'book_id', l.book_id,
				'return_book', to_char(l.return_book,'YYYY-MM-DD"T"HH24:MI:SS".00000Z"'),
				'created_at', to_char(l.created_at, 'YYYY-MM-DD"T"HH24:MI:SS".00000Z"'),
				'updated_at', to_char(l.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS".00000Z"')
			))
			FROM lend_books l
			WHERE l.book_id = b.id
		) AS lendH
	FROM books b
	ORDER BY b.id DESC
	`

	rows, err := br.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookV2

	for rows.Next() {
		var book models.BookV2
		var lendString *string

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.LiteraryGenre,
			&book.CreatedAt,
			&lendString,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los resultados: %v", err)
		}
		if lendString != nil {
			var lendBooks []*models.LendBookStr
			err = json.Unmarshal([]byte(*lendString), &lendBooks)
			if err != nil {
				utils.LogErrorData("errorMarshall", "errorMarshall", lendString)
				return nil, err
			}
			book.LendHistory = lendBooks
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar por los resultados: %v", err)
	}

	return books, nil
}
