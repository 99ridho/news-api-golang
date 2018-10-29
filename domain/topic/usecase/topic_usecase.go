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

	pagination.Limit = limit
	if resultLength > 0 && resultLength == int(limit) {
		pagination.NextCursor = result[resultLength-1].ID
	}

	return result, pagination, nil
}

func (uc *topicUseCaseImplementation) InsertTopic(ctx context.Context, topic *models.Topic) (*models.Topic, error) {
	_, err := uc.repo.Store(ctx, topic)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (uc *topicUseCaseImplementation) UpdateTopic(ctx context.Context, topic *models.Topic) (*models.Topic, error) {
	return uc.repo.Update(ctx, topic)
}

func (uc *topicUseCaseImplementation) DeleteTopic(ctx context.Context, id int64) (bool, error) {
	return uc.repo.Delete(ctx, id)
}
