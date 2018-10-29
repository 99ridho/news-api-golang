package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// News is a model that represent a news object
type News struct {
	ID                     int64          `json:"id" db:"id"`
	Author                 string         `json:"author" db:"author"`
	Slug                   string         `json:"slug" db:"slug"`
	Title                  string         `json:"title" db:"title"`
	Description            string         `json:"description" db:"description"`
	Content                string         `json:"content" db:"content"`
	TopicIDs               []int64        `json:"-"`
	Status                 string         `json:"status" db:"status"`
	PublishedAtNullableSQL mysql.NullTime `json:"-" db:"published_at"`
	PublishedAt            time.Time      `json:"published_at"`
	CreatedAt              time.Time      `json:"-" db:"created_at"`
	UpdatedAt              mysql.NullTime `json:"-" db:"updated_at"`
}
