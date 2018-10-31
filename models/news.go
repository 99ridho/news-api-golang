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
	Status                 NewsStatus     `json:"status" db:"status"`
	PublishedAtNullableSQL mysql.NullTime `json:"-" db:"published_at"`
	PublishedAt            time.Time      `json:"published_at"`
	CreatedAt              time.Time      `json:"-" db:"created_at"`
	UpdatedAt              mysql.NullTime `json:"-" db:"updated_at"`
}

func (n *News) MarkPublished() {
	n.Status = NewsStatusPublished
	n.PublishedAt = time.Now()
	n.PublishedAtNullableSQL = mysql.NullTime{
		Time:  n.PublishedAt,
		Valid: true,
	}
}

func (n *News) MarkDrafted() {
	n.Status = NewsStatusDraft
	n.PublishedAt = time.Time{}
	n.PublishedAtNullableSQL = mysql.NullTime{
		Time:  n.PublishedAt,
		Valid: false,
	}
}

func (n *News) MarkDeleted() {
	n.Status = NewsStatusDeleted
	n.PublishedAt = time.Time{}
	n.PublishedAtNullableSQL = mysql.NullTime{
		Time:  n.PublishedAt,
		Valid: false,
	}
}
