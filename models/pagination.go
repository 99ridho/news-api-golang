package models

// Pagination is a representation of data pagination (previous, next, and limit)
type Pagination struct {
	PreviousCursor int64 `json:"prev_cursor"`
	NextCursor     int64 `json:"next_cursor"`
	Limit          int64 `json:"limit"`
}
