package newsusecase_test

import (
	"context"
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
