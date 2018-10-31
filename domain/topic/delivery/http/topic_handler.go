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

type TopicMutationHandler func(ctx context.Context, t *models.Topic) (*models.Topic, error)

type TopicHandler struct {
	UseCase topic.TopicUseCase
}

func (h *TopicHandler) mutateTopic(c echo.Context, mutationHandler TopicMutationHandler) error {
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

	result, err := mutationHandler(ctx, req.Topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Insert topic failed").Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(http.StatusOK, &models.GeneralResponse{
		Data: &TopicMutationResponse{
			Topic: result,
		},
		ErrorMessage: "",
		Message:      "OK",
	})
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
	return h.mutateTopic(c, func(ctx context.Context, t *models.Topic) (*models.Topic, error) {
		return h.UseCase.InsertTopic(ctx, t)
	})
}

func (h *TopicHandler) UpdateTopic(c echo.Context) error {
	return h.mutateTopic(c, func(ctx context.Context, t *models.Topic) (*models.Topic, error) {
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.Wrap(err, "Topic ID must int")
		}

		t.ID = int64(intId)
		return h.UseCase.UpdateTopic(ctx, t)
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
