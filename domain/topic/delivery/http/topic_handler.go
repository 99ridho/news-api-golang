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
	req := new(FetchTopicRequest)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Request data invalid").Error(),
			Message:      "Fail",
		})
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	topics, pagination, err := h.UseCase.FetchTopics(ctx, req.Limit, req.NextCursor)
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
	req := new(MutateTopicRequest)
	err := c.Bind(req)
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

	result, err := h.UseCase.InsertTopic(ctx, req.Topic)
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

	req := new(MutateTopicRequest)
	err = c.Bind(req)
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

	req.ID = int64(intId)
	result, err := h.UseCase.UpdateTopic(ctx, req.Topic)
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

func (h *TopicHandler) DeleteTopic(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Topic ID must int").Error(),
			Message:      "Fail",
		})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ok, err := h.UseCase.DeleteTopic(ctx, int64(intId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Delete topic failed").Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(http.StatusOK, &models.GeneralResponse{
		Data: &DeleteTopicResponse{
			IsSuccess: ok,
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
	g.DELETE("/:id", handler.DeleteTopic)
}
