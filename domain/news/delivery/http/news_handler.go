package newshttpdelivery

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"gitlab.com/99ridho/news-api/domain/news"
	"gitlab.com/99ridho/news-api/models"
)

type NewsHandler struct {
	usecase news.NewsUseCase
}

func (h *NewsHandler) FetchNews(c echo.Context) error {
	params := new(FetchNewsRequest)
	if err := c.Bind(params); err != nil {
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

	fetchParams := &models.FetchNewsParam{
		Pagination: &models.Pagination{Limit: params.Limit, NextCursor: params.NextCursor},
		Status:     params.Status,
	}

	topicIDs := make([]int64, 0)
	if params.Topic != "" {
		ids := strings.Split(params.Topic, ",")
		for _, strID := range ids {
			num, err := strconv.Atoi(strID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, &models.GeneralResponse{
					Data:         nil,
					ErrorMessage: errors.Wrap(err, "Request data invalid").Error(),
					Message:      "Fail",
				})
			}
			topicIDs = append(topicIDs, int64(num))
		}
	}
	fetchParams.TopicIDs = topicIDs

	result, pagination, err := h.usecase.FetchNewsByParams(ctx, fetchParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: err.Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(200, &models.GeneralResponse{
		Data: &NewsResponse{
			News:       result,
			Pagination: pagination,
		},
		ErrorMessage: "",
		Message:      "OK",
	})
}

func (h *NewsHandler) DeleteNews(c echo.Context) error {
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

	ok, err := h.usecase.DeleteNews(ctx, int64(intId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.GeneralResponse{
			Data:         nil,
			ErrorMessage: errors.Wrap(err, "Delete topic failed").Error(),
			Message:      "Fail",
		})
	}

	return c.JSON(http.StatusOK, &models.GeneralResponse{
		Data: &DeleteNewsResponse{
			IsSuccess: ok,
		},
		ErrorMessage: "",
		Message:      "OK",
	})
}

func InitializeNewsHandler(r *echo.Echo, usecase news.NewsUseCase) {
	handler := &NewsHandler{usecase}

	g := r.Group("/news")

	g.GET("", handler.FetchNews)
	g.DELETE("/:id", handler.DeleteNews)
}
