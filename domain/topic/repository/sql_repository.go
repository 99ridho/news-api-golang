package topicrepository

import (
	"context"
	"errors"

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
	rows, err := repo.Conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
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
	panic("not implemented")
}

func (repo *topicSQLRepository) Update(ctx context.Context, topic *models.Topic) (*models.Topic, error) {
	panic("not implemented")
}

func (repo *topicSQLRepository) Delete(ctx context.Context, id int64) (bool, error) {
	panic("not implemented")
}
