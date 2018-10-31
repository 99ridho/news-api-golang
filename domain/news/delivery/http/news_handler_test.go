package newshttpdelivery_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/99ridho/news-api/domain/news/delivery/http"
	"gitlab.com/99ridho/news-api/domain/news/mocks"
	"gitlab.com/99ridho/news-api/models"
)

func TestFetchNewsHandler(t *testing.T) {
	mockUseCase := new(newsmocks.NewsUseCase)
	news := make([]*models.News, 0)
	pagination := &models.Pagination{
		Limit:      10,
		NextCursor: 0,
	}

	mockUseCase.On("FetchNewsByParams", mock.Anything, mock.AnythingOfType("*models.FetchNewsParam")).
		Return(news, pagination, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/news", nil)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := &newshttpdelivery.NewsHandler{
		UseCase: mockUseCase,
	}

	handler.FetchNews(ctx)
	response := new(models.GeneralResponse)

	mappingError := json.NewDecoder(rec.Result().Body).Decode(response)
	assert.NoError(t, mappingError)
	fmt.Println(response)

	data := response.Data.(map[string]interface{})
	assert.Equal(t, float64(10), data["pagination"].(map[string]interface{})["limit"].(float64))
	assert.Equal(t, float64(0), data["pagination"].(map[string]interface{})["next_cursor"].(float64))
}
