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

func InitializeTopicHandler(r *echo.Echo, usecase topic.TopicUseCase) {
	handler := &TopicHandler{
		UseCase: usecase,
	}

	g := r.Group("/topic")

	g.GET("", handler.FetchTopics)
}
