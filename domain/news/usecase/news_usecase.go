package newsusecase

import (
	"context"

	"gitlab.com/99ridho/news-api/domain/news"
	"gitlab.com/99ridho/news-api/models"
)

type newsUseCaseImplementation struct {
	repo news.NewsRepository
}

func NewNewsUseCaseImplementation(repo news.NewsRepository) news.NewsUseCase {
	return &newsUseCaseImplementation{repo}
}

func (uc *newsUseCaseImplementation) FetchNewsByParams(ctx context.Context, params *models.FetchNewsParam) ([]*models.News, *models.Pagination, error) {
	if params.Limit == 0 {
		params.Limit = 10
	}

	result, err := uc.repo.FetchByParams(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	pagination := new(models.Pagination)
	resultLength := len(result)

	pagination.Limit = params.Limit
	if resultLength > 0 && resultLength == int(params.Limit) {
		pagination.NextCursor = result[resultLength-1].ID
	}

	return result, pagination, nil
}

func (uc *newsUseCaseImplementation) InsertNews(ctx context.Context, news *models.News) (*models.News, error) {
	panic("not implemented")
}

func (uc *newsUseCaseImplementation) UpdateNews(ctx context.Context, news *models.News) (*models.News, error) {
	panic("not implemented")
}

func (uc *newsUseCaseImplementation) DeleteNews(ctx context.Context, id int64) (bool, error) {
	return uc.repo.Delete(ctx, id)
}
