package newshttpdelivery

import (
	"gitlab.com/99ridho/news-api/models"
)

type FetchNewsRequest struct {
	Topic      string            `query:"topic"`
	Limit      int64             `query:"limit"`
	NextCursor int64             `query:"next_cursor"`
	Status     models.NewsStatus `query:"status"`
}

type MutateNewsRequest struct {
	News      *models.News `json:"news"`
	NewsTopic []int64      `json:"topics"`
}
