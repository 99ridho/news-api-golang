package topicusecase_test

import (
	"context"
	"errors"
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
	assert.Equal(t, int64(0), pagination.NextCursor)
	assert.Len(t, result, len(result))
}

func TestSuccessInsertTopic(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)
	topic := &models.Topic{
		ID:   3,
		Name: "Moto GP",
		Slug: "moto-gp",
	}

	repoMock.On("Store", mock.Anything, mock.AnythingOfType("*models.Topic")).Return(int64(3), nil)

	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)
	result, err := uc.InsertTopic(context.TODO(), topic)

	assert.NoError(t, err)
	assert.Equal(t, topic, result)
}

func TestFailedInsertTopic(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)
	topic := &models.Topic{
		ID:   3,
		Name: "Moto GP",
		Slug: "moto-gp",
	}

	repoMock.On("Store", mock.Anything, mock.AnythingOfType("*models.Topic")).Return(int64(0), errors.New("Failed"))

	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)
	result, err := uc.InsertTopic(context.TODO(), topic)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSuccessUpdateTopic(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)
	topic := &models.Topic{
		ID:   3,
		Name: "Moto GP",
		Slug: "moto-gp",
	}

	repoMock.On("Update", mock.Anything, mock.AnythingOfType("*models.Topic")).Return(topic, nil)

	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)
	result, err := uc.UpdateTopic(context.TODO(), topic)

	assert.NoError(t, err)
	assert.Equal(t, topic, result)
}

func TestFailedUpdateTopic(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)
	topic := &models.Topic{
		ID:   3,
		Name: "Moto GP",
		Slug: "moto-gp",
	}

	repoMock.On("Update", mock.Anything, mock.AnythingOfType("*models.Topic")).Return(nil, errors.New("Failed"))

	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)
	result, err := uc.UpdateTopic(context.TODO(), topic)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSuccessDeleteTopic(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)

	repoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(true, nil)

	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)
	result, err := uc.DeleteTopic(context.TODO(), 1)

	assert.NoError(t, err)
	assert.True(t, result)
}

func TestFailedDeleteTopic(t *testing.T) {
	repoMock := new(topicmocks.TopicRepository)

	repoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(false, errors.New("failed"))

	uc := topicusecase.NewTopicUseCaseImplementation(repoMock)
	result, err := uc.DeleteTopic(context.TODO(), 1)

	assert.Error(t, err)
	assert.False(t, result)
}
