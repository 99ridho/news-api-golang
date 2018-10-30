package models

// Pagination is a representation of data pagination (previous, next, and limit)
type Pagination struct {
	NextCursor int64 `json:"next_cursor" query:"next_cursor"`
	Limit      int64 `json:"limit" query:"limit"`
}
