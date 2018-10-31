package topichttpdelivery

import "gitlab.com/99ridho/news-api/models"

type FetchTopicResponse struct {
	Topics     []*models.Topic    `json:"topics"`
	Pagination *models.Pagination `json:"pagination"`
}

type TopicMutationResponse struct {
	Topic *models.Topic `json:"topic"`
}

type DeleteTopicResponse struct {
	IsSuccess bool `json:"is_success"`
}
