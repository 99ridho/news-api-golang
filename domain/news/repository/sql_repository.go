package newsrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gitlab.com/99ridho/news-api/domain/news"
	"gitlab.com/99ridho/news-api/models"
)

type newsSQLRepository struct {
	Conn *sqlx.DB
}

func NewNewsSQLRepository(conn *sqlx.DB) news.NewsRepository {
	return &newsSQLRepository{conn}
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

func (repo *newsSQLRepository) transaction(ctx context.Context, handler func(tx *sqlx.Tx) error) (err error) {
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
	return err
}

func (repo *newsSQLRepository) FetchByParams(ctx context.Context, params *models.FetchNewsParam) ([]*models.News, error) {
	queryArgs := make([]interface{}, 0)

	sq := squirrel.Select("n.id", "n.author", "n.slug", "n.title", "n.description", "n.content", "n.status", "n.published_at", "n.created_at", "n.updated_at").
		From("news n").
		Where("n.id > ?", params.Pagination.NextCursor).
		Limit(uint64(params.Limit))

	queryArgs = append(queryArgs, params.Pagination.NextCursor)

	if params.Status != "" {
		sq = sq.Where("n.status = ?", params.Status)
		queryArgs = append(queryArgs, params.Status)
	}

	if len(params.TopicIDs) > 0 {
		sq = sq.Join("news_topic nt ON n.id = nt.news_id").
			Where(squirrel.Eq{"nt.topic_id": params.TopicIDs}).
			GroupBy("n.id")

		for _, topicID := range params.TopicIDs {
			queryArgs = append(queryArgs, topicID)
		}
	}

	query, _, err := sq.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Can't build query")
	}

	return repo.fetch(ctx, query, queryArgs...)
}

func (repo *newsSQLRepository) FetchById(ctx context.Context, id int64) (*models.News, error) {
	query := "SELECT * FROM `news` WHERE `id` = ?"
	return repo.fetchSingle(ctx, query, id)
}

func (repo *newsSQLRepository) FetchBySlug(ctx context.Context, slug string) (*models.News, error) {
	query := "SELECT * FROM `news` WHERE `slug` = ?"
	return repo.fetchSingle(ctx, query, slug)
}

func (repo *newsSQLRepository) FetchByStatus(ctx context.Context, status models.NewsStatus) ([]*models.News, error) {
	query := "SELECT * FROM `news` WHERE `status` = ?"
	return repo.fetch(ctx, query, status)
}

func (repo *newsSQLRepository) Store(ctx context.Context, news *models.News) (int64, error) {
	newsID := int64(0)

	insertNewsError := repo.transaction(ctx, func(tx *sqlx.Tx) error {
		insertNewsQuery := "INSERT INTO `news` (`author`, `slug`, `title`, `description`, `content`, `status`) VALUES (?,?,?,?,?,?)"
		stmt, err := tx.PreparexContext(ctx, insertNewsQuery)
		if err != nil {
			return errors.Wrap(err, "Prepare statement failed")
		}

		insertNewsResult, err := stmt.ExecContext(ctx, news.Author, news.Slug, news.Title, news.Description, news.Content, news.Status)
		if err != nil {
			return errors.Wrap(err, "Can't insert news")
		}

		id, err := insertNewsResult.LastInsertId()
		if err != nil {
			return errors.Wrap(err, "Insert news result have an error")
		}

		insertNewsTopicQueryBuilder := squirrel.Insert("news_topic").Columns("news_id", "topic_id")
		insertNewsTopicArgs := make([]interface{}, 0)
		for _, topicID := range news.TopicIDs {
			insertNewsTopicQueryBuilder = insertNewsTopicQueryBuilder.Values(newsID, topicID)
			insertNewsTopicArgs = append(insertNewsTopicArgs, newsID)
			insertNewsTopicArgs = append(insertNewsTopicArgs, topicID)
		}

		insertNewsTopicQuery, _, err := insertNewsTopicQueryBuilder.ToSql()
		if err != nil {
			return errors.Wrap(err, "Can't build insert news topic query")
		}

		stmt, err = tx.PreparexContext(ctx, insertNewsTopicQuery)
		if err != nil {
			return errors.Wrap(err, "Prepare statement failed")
		}

		_, err = stmt.ExecContext(ctx, insertNewsTopicArgs...)
		if err != nil {
			return errors.Wrap(err, "Can't insert news topic")
		}

		newsID = id
		return nil
	})

	return newsID, insertNewsError
}

func (repo *newsSQLRepository) Update(ctx context.Context, news *models.News) (*models.News, error) {
	updateNewsError := repo.transaction(ctx, func(tx *sqlx.Tx) error {
		updateQuery := "UPDATE `news` SET `author` = ?, `title` = ?, `description` = ?, `content` = ?, `status` = ?, `published_at` = ?, `updated_at` = ? WHERE `id` = ?"
		stmt, err := tx.PreparexContext(ctx, updateQuery)
		if err != nil {
			news = nil
			return errors.Wrap(err, "Prepare statement failed")
		}

		updateResult, err := stmt.ExecContext(ctx, news.Author, news.Title, news.Description, news.Content, news.Status, news.PublishedAtNullableSQL, time.Now().Format("2006-01-02 15:04:05"), news.ID)
		if err != nil {
			news = nil
			return errors.Wrap(err, "Can't update news")
		}

		rowsAffected, err := updateResult.RowsAffected()
		if err != nil {
			news = nil
			errors.Wrap(err, "Can't check updated rows id")
		}
		if rowsAffected != 1 {
			news = nil
			return errors.Wrap(err, "Weird behavior, rows affected more 1")
		}

		if len(news.TopicIDs) > 0 {
			// delete news topic first
			deleteTopicQuery := "DELETE FROM `news_topic` WHERE `news_id` = ?"
			stmt, err := tx.PreparexContext(ctx, deleteTopicQuery)
			if err != nil {
				news = nil
				return errors.Wrap(err, "Prepare statement failed")
			}

			_, err = stmt.ExecContext(ctx, news.ID)
			if err != nil {
				news = nil
				return errors.Wrap(err, "Can't update news topic")
			}

			insertNewsTopicQueryBuilder := squirrel.Insert("news_topic").Columns("news_id", "topic_id")
			insertNewsTopicArgs := make([]interface{}, 0)
			for _, topicID := range news.TopicIDs {
				insertNewsTopicQueryBuilder = insertNewsTopicQueryBuilder.Values(news.ID, topicID)
				insertNewsTopicArgs = append(insertNewsTopicArgs, news.ID)
				insertNewsTopicArgs = append(insertNewsTopicArgs, topicID)
			}

			insertNewsTopicQuery, _, err := insertNewsTopicQueryBuilder.ToSql()
			if err != nil {
				return errors.Wrap(err, "Can't build insert news topic query")
			}

			stmt, err = tx.PreparexContext(ctx, insertNewsTopicQuery)
			if err != nil {
				return errors.Wrap(err, "Prepare statement failed")
			}

			_, err = stmt.ExecContext(ctx, insertNewsTopicArgs...)
			if err != nil {
				return errors.Wrap(err, "Can't insert news topic")
			}
		}

		return nil
	})

	return news, updateNewsError
}

func (repo *newsSQLRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM `news` WHERE `id` = ?"

	stmt, err := repo.Conn.PreparexContext(ctx, query)
	if err != nil {
		return false, errors.Wrap(err, "Prepare statement failed")
	}

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return false, errors.Wrap(err, "Deleting news failed")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "Can't check deleted rows id")
	}
	if rowsAffected != 1 {
		return false, fmt.Errorf("Weird behavior, row affected : %d", rowsAffected)
	}

	return true, nil
}
