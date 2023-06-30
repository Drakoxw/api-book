package models

import "time"

type User struct {
	Id          int64       `json:"id,omitempty"`
	Username    string      `json:"username"`
	Password    string      `json:"password"`
	Email       string      `json:"email"`
	CreatedAt   time.Time   `json:"create_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	LendHistory []*LendBook `json:"LendHistory,omitempty"`
}

type UpdateUser struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
