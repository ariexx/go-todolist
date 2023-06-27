package model

import "time"

type Todo struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
