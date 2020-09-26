package models

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  []byte    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
