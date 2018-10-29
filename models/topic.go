package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// Topic is a model that represent a topic object
type Topic struct {
	ID        int64          `json:"id" db:"id"`
	Slug      string         `json:"slug" db:"slug"`
	Name      string         `json:"name" db:"name"`
	CreatedAt time.Time      `json:"-" db:"created_at"`
	UpdatedAt mysql.NullTime `json:"-" db:"updated_at"`
}
