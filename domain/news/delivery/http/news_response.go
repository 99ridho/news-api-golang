package newshttpdelivery

import (
	"gitlab.com/99ridho/news-api/models"
)

type NewsResponse struct {
	News       []*models.News     `json:"news"`
	Pagination *models.Pagination `json:"pagination"`
}
