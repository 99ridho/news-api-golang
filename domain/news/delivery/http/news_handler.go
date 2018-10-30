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
	params := new(models.FetchNewsParam)
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

	if params.Pagination == nil {
		params.Pagination = &models.Pagination{Limit: 10, NextCursor: 0}
	}

	topicIDs := make([]int64, 0)
	if params.TopicIDQuery != "" {
		ids := strings.Split(params.TopicIDQuery, ",")
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
	params.TopicIDs = topicIDs

	result, pagination, err := h.usecase.FetchNewsByParams(ctx, params)
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

func InitializeNewsHandler(r *echo.Echo, usecase news.NewsUseCase) {
	handler := &NewsHandler{usecase}

	g := r.Group("/news")

	g.GET("", handler.FetchNews)
}
