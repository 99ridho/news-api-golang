package topic

import (
	"context"

	"gitlab.com/99ridho/news-api/models"
)

// TopicUseCase is representing a topic business logic
type TopicUseCase interface {
	FetchTopics(ctx context.Context, limit, cursor int64) ([]*models.Topic, error)
	InsertTopic(ctx context.Context, topic *models.Topic) (*models.Topic, error)
	UpdateTopic(ctx context.Context, topic *models.Topic) (*models.Topic, error)
	DeleteTopic(ctx context.Context, topic *models.Topic) (bool, error)
}
