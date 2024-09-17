package models

import "time"

type User struct {
	ID        int
	LastName  string
	FirstName string
	Tel       string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
