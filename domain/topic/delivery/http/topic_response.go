package topichttpdelivery

import "gitlab.com/99ridho/news-api/models"

type FetchTopicResponse struct {
	Topics     []*models.Topic    `json:"topics"`
	Pagination *models.Pagination `json:"pagination"`
}

type InsertTopicResponse struct {
	Topic *models.Topic `json:"inserted_topic"`
}

type UpdateTopicResponse struct {
	Topic *models.Topic `json:"updated_topic"`
}

type DeleteTopicResponse struct {
	IsSuccess bool `json:"is_success"`
}
