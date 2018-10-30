package topichttpdelivery

import (
	"gitlab.com/99ridho/news-api/models"
)

type FetchTopicRequest struct {
	Limit      int64 `query:"limit"`
	NextCursor int64 `query:"next_cursor"`
}

type MutateTopicRequest struct {
	*models.Topic
}
