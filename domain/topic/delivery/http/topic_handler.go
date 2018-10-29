package topichttpdelivery

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/labstack/echo"
	"gitlab.com/99ridho/news-api/domain/topic"
	"gitlab.com/99ridho/news-api/models"
)

type TopicHandler struct {
	UseCase topic.TopicUseCase
}

func (h *TopicHandler) FetchTopics(c echo.Context) error {
	limit := int64(0)
	cursor := int64(0)
	query := c.Request().URL.Query()
	limitQuery := query.Get("limit")
	cursorQuery := query.Get("cursor")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if limitQuery != "" {
		limitInt, err := strconv.Atoi(limitQuery)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
				Data:         nil,
				ErrorMessage: errors.Wrap(err, "Can't convert cursor to int").Error(),
				Message:      "Fail",
			})
		}
		limit = int64(limitInt)
	}
	if cursorQuery != "" {
		cursorInt, err := strconv.Atoi(cursorQuery)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
				Data:         nil,
				ErrorMessage: errors.Wrap(err, "Can't convert cursor to int").Error(),
				Message:      "Fail",
			})
		}
		cursor = int64(cursorInt)
	}

	topics, pagination, err := h.UseCase.FetchTopics(ctx, limit, cursor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Can't fetch topics").Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(http.StatusOK, &models.GeneralResponse{
		Data: &FetchTopicResponse{
			Topics:     topics,
			Pagination: pagination,
		},
		ErrorMessage: "",
		Message:      "OK",
	})
}

func (h *TopicHandler) InsertTopic(c echo.Context) error {
	topic := new(models.Topic)
	err := c.Bind(topic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Request data invalid").Error(),
			Message:      "Fail",
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UseCase.InsertTopic(ctx, topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Insert topic failed").Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(http.StatusOK, &models.GeneralResponse{
		Data: &InsertTopicResponse{
			Topic: result,
		},
		ErrorMessage: "",
		Message:      "OK",
	})
}

func (h *TopicHandler) UpdateTopic(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Topic ID must int").Error(),
			Message:      "Fail",
		})
	}

	topic := new(models.Topic)
	err = c.Bind(topic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Request data invalid").Error(),
			Message:      "Fail",
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	topic.ID = int64(intId)
	result, err := h.UseCase.UpdateTopic(ctx, topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Update topic failed").Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(http.StatusOK, &models.GeneralResponse{
		Data: &UpdateTopicResponse{
			Topic: result,
		},
		ErrorMessage: "",
		Message:      "OK",
	})
}

func InitializeTopicHandler(r *echo.Echo, usecase topic.TopicUseCase) {
	handler := &TopicHandler{
		UseCase: usecase,
	}

	g := r.Group("/topic")

	g.GET("", handler.FetchTopics)
	g.POST("", handler.InsertTopic)
	g.PUT("/:id", handler.UpdateTopic)
}
