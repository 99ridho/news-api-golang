package topichttpdelivery

import (
	"gitlab.com/99ridho/news-api/models"
)

type FetchTopicRequest struct {
	Limit      string `query:"limit"`
	NextCursor string `query:"next_cursor"`
}

type MutateTopicRequest struct {
	*models.Topic
}
