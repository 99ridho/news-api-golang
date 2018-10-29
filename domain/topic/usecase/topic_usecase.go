package topicusecase

import (
	"context"

	"gitlab.com/99ridho/news-api/domain/topic"
	"gitlab.com/99ridho/news-api/models"
)

type topicUseCaseImplementation struct {
	repo topic.TopicRepository
}

func NewTopicUseCaseImplementation(repo topic.TopicRepository) topic.TopicUseCase {
	return &topicUseCaseImplementation{repo}
}

func (uc *topicUseCaseImplementation) FetchTopics(ctx context.Context, limit int64, cursor int64) ([]*models.Topic, *models.Pagination, error) {
	if limit == 0 {
		limit = 10
	}

	result, err := uc.repo.Fetch(ctx, cursor, limit)
	if err != nil {
		return nil, nil, err
	}

	pagination := new(models.Pagination)
	resultLength := len(result)

	if resultLength > 0 {
		pagination.Limit = limit
		pagination.NextCursor = result[resultLength-1].ID
	}

	return result, pagination, nil
}

func (uc *topicUseCaseImplementation) InsertTopic(ctx context.Context, topic *models.Topic) (*models.Topic, error) {
	panic("not implemented")
}

func (uc *topicUseCaseImplementation) UpdateTopic(ctx context.Context, topic *models.Topic) (*models.Topic, error) {
	panic("not implemented")
}

func (uc *topicUseCaseImplementation) DeleteTopic(ctx context.Context, topic *models.Topic) (bool, error) {
	panic("not implemented")
}
