package topichttpdelivery

import "gitlab.com/99ridho/news-api/models"

type FetchTopicResponse struct {
	Topics     []*models.Topic    `json:"topics"`
	Pagination *models.Pagination `json:"pagination"`
}
