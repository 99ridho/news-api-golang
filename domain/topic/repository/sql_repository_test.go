package topicrepository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
	"gitlab.com/99ridho/news-api/domain/topic/repository"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetchTopics(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := topicrepository.NewTopicSQLRepository(sqlxMockDb)
	query := "SELECT \\* FROM \\`topic\\`"
	rows := sqlmock.NewRows([]string{"id", "slug", "name", "created_at", "updated_at"}).
		AddRow(1, "motogp", "MotoGP", time.Now(), time.Now()).
		AddRow(2, "fifawc", "World Cup 2018", time.Now(), time.Now())

	mock.ExpectQuery(query).WillReturnRows(rows)

	topics, err := repo.Fetch(context.Background(), 0, 6)

	assert.NoError(t, err)
	assert.Len(t, topics, 2)
}

func TestFetchTopicByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := topicrepository.NewTopicSQLRepository(sqlxMockDb)
	query := "SELECT \\* FROM \\`topic\\`"
	rows := sqlmock.NewRows([]string{"id", "slug", "name", "created_at", "updated_at"}).
		AddRow(1, "motogp", "MotoGP", time.Now(), time.Now()).
		AddRow(2, "fifawc", "World Cup 2018", time.Now(), time.Now())

	mock.ExpectQuery(query).WillReturnRows(rows)

	topic, err := repo.FetchById(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "motogp", topic.Slug)
}

func TestFetchTopicBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := topicrepository.NewTopicSQLRepository(sqlxMockDb)
	query := "SELECT \\* FROM \\`topic\\`"
	rows := sqlmock.NewRows([]string{"id", "slug", "name", "created_at", "updated_at"}).
		AddRow(1, "motogp", "MotoGP", time.Now(), time.Now()).
		AddRow(2, "fifawc", "World Cup 2018", time.Now(), time.Now())

	mock.ExpectQuery(query).WillReturnRows(rows)

	topic, err := repo.FetchBySlug(context.Background(), "motogp")

	assert.NoError(t, err)
	assert.Equal(t, "motogp", topic.Slug)
}
