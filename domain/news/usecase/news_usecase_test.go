package newsusecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/99ridho/news-api/models"

	"gitlab.com/99ridho/news-api/domain/news/mocks"
	"gitlab.com/99ridho/news-api/domain/news/usecase"
)

func TestFetchNewsByParams(t *testing.T) {
	repoMock := new(newsmocks.NewsRepository)
	params := &models.FetchNewsParam{
		TopicIDs: []int64{},
		Status:   models.NewsStatusPublished,
		Pagination: &models.Pagination{
			Limit:      10,
			NextCursor: 0,
		},
	}
	mockedResult := []*models.News{
		&models.News{
			ID:    1,
			Title: "Sapijan",
		},
		&models.News{
			ID:    2,
			Title: "Sapijan",
		},
		&models.News{
			ID:    3,
			Title: "Sapijan",
		},
		&models.News{
			ID:    4,
			Title: "Sapijan",
		},
		&models.News{
			ID:    5,
			Title: "Sapijan",
		},
	}

	repoMock.
		On("FetchByParams", mock.Anything, mock.AnythingOfType("*models.FetchNewsParam")).
		Return(mockedResult, nil)

	uc := newsusecase.NewNewsUseCaseImplementation(repoMock)
	result, pagination, err := uc.FetchNewsByParams(context.TODO(), params)

	assert.NoError(t, err)
	assert.Equal(t, int64(0), pagination.NextCursor)
	assert.Len(t, result, len(result))
}

func TestSuccessDeleteNews(t *testing.T) {
	repoMock := new(newsmocks.NewsRepository)

	repoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(true, nil)

	uc := newsusecase.NewNewsUseCaseImplementation(repoMock)
	result, err := uc.DeleteNews(context.TODO(), 1)

	assert.NoError(t, err)
	assert.True(t, result)
}

func TestFailedDeleteNews(t *testing.T) {
	repoMock := new(newsmocks.NewsRepository)

	repoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(false, errors.New("fail"))

	uc := newsusecase.NewNewsUseCaseImplementation(repoMock)
	result, err := uc.DeleteNews(context.TODO(), 1)

	assert.Error(t, err)
	assert.False(t, result)
}

func TestSuccessInsertNews(t *testing.T) {
	repoMock := new(newsmocks.NewsRepository)
	repoMock.On("Store", mock.Anything, mock.AnythingOfType("*models.News")).Return(int64(1), nil)

	n := &models.News{
		ID:   1,
		Slug: "halo",
	}

	uc := newsusecase.NewNewsUseCaseImplementation(repoMock)
	result, err := uc.InsertNews(context.TODO(), n)

	assert.NoError(t, err)
	assert.Equal(t, n, result)
}

func TestFailInsertNews(t *testing.T) {
	repoMock := new(newsmocks.NewsRepository)
	repoMock.On("Store", mock.Anything, mock.AnythingOfType("*models.News")).Return(int64(0), errors.New("fail"))

	n := &models.News{
		ID:   1,
		Slug: "halo",
	}

	uc := newsusecase.NewNewsUseCaseImplementation(repoMock)
	result, err := uc.InsertNews(context.TODO(), n)

	assert.Error(t, err)
	assert.Nil(t, result)
}
