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

var CreateDataUser = `
INSERT INTO users (username, password, email, created_at)
VALUES
	('user2', 'password222', 'user2@example.com', NOW()),
	('user3', 'password333', 'user3@example.com', NOW()),
	('user4', 'password444', 'user4@example.com', NOW());
	('user5', 'password555', 'user5@example.com', NOW()),

`
