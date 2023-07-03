package models

import "time"

type Book struct {
	Id            int64       `json:"id,omitempty"`
	Title         string      `json:"title"`
	Author        string      `json:"author"`
	LiteraryGenre string      `json:"literary_genre"`
	CreatedAt     time.Time   `json:"created_at,omitempty"`
	LendHistory   []*LendBook `json:"lendHistory,omitempty"`
}
type BookV2 struct {
	Id            int64          `json:"id"`
	Title         string         `json:"title"`
	Author        string         `json:"author"`
	LiteraryGenre string         `json:"literary_genre"`
	CreatedAt     time.Time      `json:"created_at"`
	LendHistory   []*LendBookStr `json:"lendHistory,omitempty"`
}

var CreateDataBook = `
INSERT INTO books (title, author, literary_genre, created_at)
VALUES
    ('Book 1', 'Author 1', 'Fantasy', NOW()),
    ('Book 2', 'Author 2', 'Romance', NOW()),
    ('Book 3', 'Author 3', 'Mystery', NOW()),
    ('Book 4', 'Author 4', 'Science Fiction', NOW()),
    ('Book 5', 'Author 5', 'Thriller', NOW()),
    ('Book 6', 'Author 1', 'Historical Fiction', NOW()),
    ('Book 7', 'Author 2', 'Horror', NOW()),
    ('Book 8', 'Author 3', 'Biography', NOW()),
    ('Book 9', 'Author 4', 'Self-Help', NOW()),
    ('Book 10', 'Author 5', 'Young Adult', NOW()),
    ('Book 11', 'Author 1', 'Drama', NOW()),
    ('Book 12', 'Author 2', 'Poety', NOW()),
    ('Book 13', 'Author 3', 'Humor', NOW()),
    ('Book 14', 'Author 4', 'Non-Fiction', NOW()),
    ('Book 15', 'Author 5', 'Classic', NOW()),
    ('Book 16', 'Author 1', 'Crime', NOW()),
    ('Book 17', 'Author 1', 'Adventure', NOW()),
    ('Book 18', 'Author 1', 'Historical', NOW()),
    ('Book 19', 'Author 2', 'Science', NOW()),
    ('Book 20', 'Author 5', 'Fiction', NOW());

`
