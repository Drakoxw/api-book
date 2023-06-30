package models

import "time"

type Book struct {
	Id            int64       `json:"id,omitempty"`
	Title         string      `json:"title"`
	Author        string      `json:"author"`
	LiteraryGenre string      `json:"literary_genre"`
	CreatedAt     time.Time   `json:"created_at,omitempty"`
	LendHistory   []*LendBook `json:"LendHistory,omitempty"`
}
