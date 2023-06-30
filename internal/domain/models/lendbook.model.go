package models

import "time"

type LendBook struct {
	Id         int64     `json:"id,omitempty"`
	UserId     int64     `json:"user_id"`
	BookId     int64     `json:"book_id"`
	ReturnBook time.Time `json:"return_book,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
