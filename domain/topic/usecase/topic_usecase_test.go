package topicusecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
	"gitlab.com/99ridho/news-api/domain/topic/mocks"
	"gitlab.com/99ridho/news-api/domain/topic/usecase"
	"gitlab.com/99ridho/news-api/models"
)

func TestFetchTopics(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)
	topics := []*models.Topic{
		&models.Topic{
			ID:   1,
			Name: "Moto GP",
			Slug: "moto-gp",
		},
		&models.Topic{
			ID:   2,
			Name: "Moto GP",
			Slug: "moto-gp",
		},
		&models.Topic{
			ID:   3,
			Name: "Moto GP",
			Slug: "moto-gp",
		},
		&models.Topic{
			ID:   4,
			Name: "Moto GP",
			Slug: "moto-gp",
		},
		&models.Topic{
			ID:   5,
			Name: "Moto GP",
			Slug: "moto-gp",
		},
	}

	repoMock.On("Fetch", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(topics, nil)
	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)

	result, pagination, err := uc.FetchTopics(context.TODO(), 10, 0)

	assert.NoError(t, err)
	assert.Equal(t, 5, pagination.NextCursor)
	assert.Len(t, result, len(result))
}
