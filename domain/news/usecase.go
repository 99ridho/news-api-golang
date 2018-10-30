package news

import (
	"context"

	"gitlab.com/99ridho/news-api/models"
)

// NewsUseCase is representing a news business logic
type NewsUseCase interface {
	FetchNewsByParams(ctx context.Context, params *models.FetchNewsParam) ([]*models.News, *models.Pagination, error)
	InsertNews(ctx context.Context, news *models.News) (*models.News, error)
	UpdateNews(ctx context.Context, news *models.News) (*models.News, error)
	DeleteNews(ctx context.Context, id int64) (bool, error)
}
