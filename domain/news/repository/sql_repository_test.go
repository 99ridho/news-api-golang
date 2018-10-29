package newsrepository_test

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gitlab.com/99ridho/news-api/domain/news/repository"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetchNewsByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := newsrepository.NewNewsSQLRepository(sqlxMockDb)
	query := "SELECT \\* FROM \\`news\\`"
	rows := sqlmock.NewRows([]string{"id", "author", "slug", "title", "description", "content", "status", "published_at", "created_at", "updated_at"}).
		AddRow(1, "fulan", "kpk-bubar", "KPK Bubar", "lorem", "lorem", "published", time.Now(), time.Now(), time.Now()).
		AddRow(2, "fulan", "lion-air-jt610-jatuh", "Lion Air JT610 jatuh", "lorem", "lorem", "published", time.Now(), time.Now(), time.Now())

	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

	news, err := repo.FetchById(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "kpk-bubar", news.Slug)
}
