package topicrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	"gitlab.com/99ridho/news-api/domain/topic"
	"gitlab.com/99ridho/news-api/models"
)

type topicSQLRepository struct {
	Conn *sqlx.DB
}

// NewTopicSQLRepository is a function to make a real implementation of TopicRepository with SQL database
func NewTopicSQLRepository(conn *sqlx.DB) topic.TopicRepository {
	return &topicSQLRepository{
		Conn: conn,
	}
}

func (repo *topicSQLRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Topic, error) {
	stmt, err := repo.Conn.PreparexContext(ctx, query)
	if err != nil {
		return []*models.Topic{}, errors.Wrap(err, "Prepare statement failed")
	}

	rows, err := stmt.QueryxContext(ctx, args...)
	defer rows.Close()
	if err != nil {
		return []*models.Topic{}, errors.Wrap(err, "Failed to fetch topic rows")
	}

	topics := make([]*models.Topic, 0)
	for rows.Next() {
		topic := new(models.Topic)
		err := rows.StructScan(topic)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}

func (repo *topicSQLRepository) fetchSingle(ctx context.Context, query string, args ...interface{}) (*models.Topic, error) {
	topics, err := repo.fetch(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if len(topics) <= 0 {
		return nil, errors.New("Topic not found")
	}

	return topics[0], nil
}

func (repo *topicSQLRepository) Fetch(ctx context.Context, cursor int64, limit int64) ([]*models.Topic, error) {
	query := "SELECT * FROM `topic` WHERE `id` > ? LIMIT ?"
	return repo.fetch(ctx, query, cursor, limit)
}

func (repo *topicSQLRepository) FetchById(ctx context.Context, id int64) (*models.Topic, error) {
	query := "SELECT * FROM `topic` WHERE `id` = ?"
	return repo.fetchSingle(ctx, query, id)
}

func (repo *topicSQLRepository) FetchBySlug(ctx context.Context, slug string) (*models.Topic, error) {
	query := "SELECT * FROM `topic` WHERE `slug` = ?"
	return repo.fetchSingle(ctx, query, slug)
}

func (repo *topicSQLRepository) Store(ctx context.Context, topic *models.Topic) (int64, error) {
	query := "INSERT INTO `topic` (`slug`, `name`) VALUES (?,?)"

	stmt, err := repo.Conn.PreparexContext(ctx, query)
	if err != nil {
		return 0, errors.Wrap(err, "Prepare statement failed")
	}

	result, err := stmt.ExecContext(ctx, topic.Slug, topic.Name)
	if err != nil {
		return 0, errors.Wrap(err, "Inserting topic failed")
	}

	return result.LastInsertId()
}

func (repo *topicSQLRepository) Update(ctx context.Context, topic *models.Topic) (*models.Topic, error) {
	sq := squirrel.Update("topic").Set("updated_at", time.Now().Format("2006-01-02 15:04:05")).Where("id = ?", topic.ID)

	updateArgs := make([]interface{}, 0)
	updateArgs = append(updateArgs, time.Now().Format("2006-01-02 15:04:05"))

	if topic.Slug != "" {
		sq = sq.Set("slug", topic.Slug)
		updateArgs = append(updateArgs, topic.Slug)
	}

	if topic.Name != "" {
		sq = sq.Set("name", topic.Name)
		updateArgs = append(updateArgs, topic.Name)
	}

	updateArgs = append(updateArgs, topic.ID)
	query, _, err := sq.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "can't build update query")
	}

	stmt, err := repo.Conn.PreparexContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "Prepare statement failed")
	}

	result, err := stmt.ExecContext(ctx, updateArgs...)
	if err != nil {
		return nil, errors.Wrap(err, "Updating topic failed")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "Can't check updated rows id")
	}
	if rowsAffected != 1 {
		return nil, fmt.Errorf("Weird behavior, row affected : %d", rowsAffected)
	}

	return repo.FetchById(ctx, topic.ID)
}

func (repo *topicSQLRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM `topic` WHERE `id` = ?"

	stmt, err := repo.Conn.PreparexContext(ctx, query)
	if err != nil {
		return false, errors.Wrap(err, "Prepare statement failed")
	}

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return false, errors.Wrap(err, "Deleting topic failed")
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
