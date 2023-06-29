package models

// import "time"

type User struct {
	Id        int64  `json:"id,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"create_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
