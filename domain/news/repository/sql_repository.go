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
	defer rows.Close()
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

func (repo *newsSQLRepository) transaction(ctx context.Context, handler func(tx *sqlx.Tx) error) {
	tx, err := repo.Conn.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = handler(tx)
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
	newsID := int64(0)
	var insertNewsError error

	repo.transaction(ctx, func(tx *sqlx.Tx) error {
		stmt, err := tx.PreparexContext(ctx, "INSERT INTO `news` (`author`, `slug`, `title`, `description`, `content`, `status`) VALUES (?,?,?,?,?,?)")
		if err != nil {
			insertNewsError = errors.Wrap(err, "Prepare statement failed")
			return insertNewsError
		}

		insertNewsResult, err := stmt.ExecContext(ctx, news.Author, news.Slug, news.Title, news.Description, news.Content, news.Status)
		if err != nil {
			insertNewsError = errors.Wrap(err, "Prepare statement failed")
			return insertNewsError
		}

		id, insertNewsError := insertNewsResult.LastInsertId()

		for _, topicID := range news.TopicIDs {
			stmt, err := tx.PreparexContext(ctx, "INSERT INTO `news_topic` (`news_id`, `topic_id`) VALUES (?,?)")
			if err != nil {
				insertNewsError = errors.Wrap(err, "Prepare statement failed")
				return insertNewsError
			}

			_, err = stmt.ExecContext(ctx, newsID, topicID)
			if err != nil {
				insertNewsError = errors.Wrap(err, "Prepare statement failed")
				return insertNewsError
			}
		}

		newsID = id
		insertNewsError = nil
		return nil
	})

	return newsID, insertNewsError
}

func (repo *newsSQLRepository) Update(ctx context.Context, news *models.News) (*models.News, error) {
	panic("not implemented")
}

func (repo *newsSQLRepository) Delete(ctx context.Context, id int64) (bool, error) {
	panic("not implemented")
}