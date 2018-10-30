package news

import (
	"context"

	"gitlab.com/99ridho/news-api/models"
)

// NewsRepository is a repository to manage a News
type NewsRepository interface {
	FetchById(ctx context.Context, id int64) (*models.News, error)
	FetchBySlug(ctx context.Context, slug string) (*models.News, error)
	FetchByStatus(ctx context.Context, status models.NewsStatus) ([]*models.News, error)
	Store(ctx context.Context, news *models.News) (int64, error)
	Update(ctx context.Context, news *models.News) (*models.News, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
