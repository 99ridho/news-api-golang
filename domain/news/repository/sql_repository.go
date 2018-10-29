package newsrepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gitlab.com/99ridho/news-api/domain/news"
	"gitlab.com/99ridho/news-api/models"
)

type newsSQLRepository struct {
	Conn *sqlx.DB
}

func (repo *newsSQLRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.News, error) {
	stmt, err := repo.Conn.PreparexContext(ctx, query)
	if err != nil {
		return []*models.News{}, errors.Wrap(err, "Prepare statement failed")
	}

	rows, err := stmt.QueryxContext(ctx, args...)
	if err != nil {
		return []*models.News{}, errors.Wrap(err, "Failed to fetch news rows")
	}

	newsList := make([]*models.News, 0)
	for rows.Next() {
		news := new(models.News)
		err := rows.StructScan(news)
		if err != nil {
			return nil, err
		}

		newsList = append(newsList, news)
	}

	return newsList, nil
}

func (repo *newsSQLRepository) fetchSingle(ctx context.Context, query string, args ...interface{}) (*models.News, error) {
	news, err := repo.fetch(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if len(news) <= 0 {
		return nil, errors.New("Topic not found")
	}

	return news[0], nil
}

func NewNewsSQLRepository(conn *sqlx.DB) news.NewsRepository {
	return &newsSQLRepository{conn}
}

func (repo *newsSQLRepository) FetchById(ctx context.Context, id int64) (*models.News, error) {
	query := "SELECT * FROM `news` WHERE `id` = ?"
	return repo.fetchSingle(ctx, query, id)
}

func (repo *newsSQLRepository) FetchBySlug(ctx context.Context, slug string) (*models.News, error) {
	query := "SELECT * FROM `news` WHERE `slug` = ?"
	return repo.fetchSingle(ctx, query, slug)
}

func (repo *newsSQLRepository) FetchByStatus(ctx context.Context, status *models.NewsStatus) ([]*models.News, error) {
	query := "SELECT * FROM `news` WHERE `status` = ?"
	return repo.fetch(ctx, query, status.String())
}

func (repo *newsSQLRepository) Store(ctx context.Context, news *models.News) (int64, error) {
	panic("not implemented")
}

func (repo *newsSQLRepository) Update(ctx context.Context, news *models.News) (*models.News, error) {
	panic("not implemented")
}

func (repo *newsSQLRepository) Delete(ctx context.Context, id int64) (bool, error) {
	panic("not implemented")
}
