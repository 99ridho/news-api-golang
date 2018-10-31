package topichttpdelivery_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	"gitlab.com/99ridho/news-api/domain/topic/delivery/http"
	"gitlab.com/99ridho/news-api/domain/topic/mocks"
	"gitlab.com/99ridho/news-api/models"

	"github.com/labstack/echo"
)

func TestFetchHandler(t *testing.T) {
	mockUseCase := new(topicmocks.TopicUseCase)
	topics := make([]*models.Topic, 0)
	pagination := &models.Pagination{
		Limit:      10,
		NextCursor: 0,
	}

	mockUseCase.On("FetchTopics", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).
		Return(topics, pagination, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/topic", nil)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := &topichttpdelivery.TopicHandler{
		UseCase: mockUseCase,
	}

	handler.FetchTopics(ctx)
	response := new(models.GeneralResponse)

	mappingError := json.NewDecoder(rec.Result().Body).Decode(response)
	assert.NoError(t, mappingError)
	fmt.Println(response)

	data := response.Data.(map[string]interface{})
	assert.Equal(t, float64(10), data["pagination"].(map[string]interface{})["limit"].(float64))
	assert.Equal(t, float64(0), data["pagination"].(map[string]interface{})["next_cursor"].(float64))
}

func TestInsertTopicHandler(t *testing.T) {
	mockUseCase := new(topicmocks.TopicUseCase)
	topic := &models.Topic{}

	mockUseCase.On("InsertTopic", mock.Anything, mock.AnythingOfType("*models.Topic")).
		Return(topic, nil)

	e := echo.New()
	payload := []byte(`
	{
		"name": "Motogp"
	}
	`)

	req := httptest.NewRequest(http.MethodPost, "/topic", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := &topichttpdelivery.TopicHandler{
		UseCase: mockUseCase,
	}

	handler.InsertTopic(ctx)
	response := new(models.GeneralResponse)

	mappingError := json.NewDecoder(rec.Result().Body).Decode(response)
	assert.NoError(t, mappingError)

	data := response.Data.(map[string]interface{})
	assert.Equal(t, "", data["topic"].(map[string]interface{})["name"].(string))
}

func TestUpdateTopicHandler(t *testing.T) {
	mockUseCase := new(topicmocks.TopicUseCase)
	topic := &models.Topic{}

	mockUseCase.On("UpdateTopic", mock.Anything, mock.AnythingOfType("*models.Topic")).
		Return(topic, nil)

	e := echo.New()
	payload := []byte(`
	{
		"name": "Motogp"
	}
	`)

	req := httptest.NewRequest(http.MethodPut, "/topic", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/topic")
	ctx.SetParamNames("id")
	ctx.SetParamValues("10")

	handler := &topichttpdelivery.TopicHandler{
		UseCase: mockUseCase,
	}

	handler.UpdateTopic(ctx)
	response := new(models.GeneralResponse)

	mappingError := json.NewDecoder(rec.Result().Body).Decode(response)
	assert.NoError(t, mappingError)

	data := response.Data.(map[string]interface{})
	assert.Equal(t, "", data["topic"].(map[string]interface{})["name"].(string))
}

func TestDeleteTopicHandler(t *testing.T) {
	mockUseCase := new(topicmocks.TopicUseCase)

	mockUseCase.On("DeleteTopic", mock.Anything, mock.AnythingOfType("int64")).
		Return(true, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/topic", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/topic")
	ctx.SetParamNames("id")
	ctx.SetParamValues("10")

	handler := &topichttpdelivery.TopicHandler{
		UseCase: mockUseCase,
	}

	handler.DeleteTopic(ctx)
	response := new(models.GeneralResponse)

	mappingError := json.NewDecoder(rec.Result().Body).Decode(response)
	assert.NoError(t, mappingError)

	data := response.Data.(map[string]interface{})
	assert.True(t, data["is_success"].(bool))
}
