package entities

import "time"

type User struct {
	ID        int       `gorm:"type:integer; primaryKey, autoIncrement, not null"`
	LastName  string    `gorm:"type:varchar(100); not null"`
	FirstName string    `gorm:"type:varchar(100); not null"`
	Tel       string    `gorm:"type:integer; not null"`
	Age       int       `gorm:"type:integer"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
