package news

import (
	"context"

	"gitlab.com/99ridho/news-api/models"
)

// NewsUseCase is representing a news business logic
type NewsUseCase interface {
	FetchNews(ctx context.Context, limit, cursor int64) ([]*models.News, error)
	InsertNews(ctx context.Context, news *models.News) (*models.News, error)
	UpdateNews(ctx context.Context, news *models.News) (*models.News, error)
	DeleteNews(ctx context.Context, news *models.News) (int64, error)
	FilterNewsByStatus(ctx context.Context, status *models.NewsStatus) ([]*models.News, error)
	FilterNewsByTopicIDs(ctx context.Context, topicIDs []int64) ([]*models.News, error)
}
