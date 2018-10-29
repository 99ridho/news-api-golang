package topic

import (
	"context"

	"gitlab.com/99ridho/news-api/models"
)

// TopicRepository is a repository to manage a Topic
type TopicRepository interface {
	FetchById(ctx context.Context, id int64) (*models.Topic, error)
	FetchBySlug(ctx context.Context, slug string) (*models.Topic, error)
	Store(ctx context.Context, topic *models.Topic) (int64, error)
	Update(ctx context.Context, topic *models.Topic) (*models.Topic, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
