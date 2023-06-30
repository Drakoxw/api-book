package repository

import (
	"api-book/internal/domain/models"
	"database/sql"
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

func (br *BookRepository) ListBooks() ([]models.Book, error) {
	query := `
		SELECT id, title, author, literary_genre, created_at
		FROM books
	`

	rows, err := br.db.Query(query)
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
