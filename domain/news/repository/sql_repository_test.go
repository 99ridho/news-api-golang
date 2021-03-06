package newsrepository_test

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gitlab.com/99ridho/news-api/domain/news/repository"
	"gitlab.com/99ridho/news-api/models"
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

func TestFetchNewsBySlug(t *testing.T) {
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

	slug := "kpk-bubar"
	news, err := repo.FetchBySlug(context.Background(), slug)

	assert.NoError(t, err)
	assert.Equal(t, slug, news.Slug)
}

func TestFetchNewsByStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := newsrepository.NewNewsSQLRepository(sqlxMockDb)
	query := "SELECT \\* FROM \\`news\\`"
	rows := sqlmock.NewRows([]string{"id", "author", "slug", "title", "description", "content", "status", "published_at", "created_at", "updated_at"}).
		AddRow(1, "fulan", "kpk-bubar", "KPK Bubar", "lorem", "lorem", "draft", time.Now(), time.Now(), time.Now()).
		AddRow(2, "fulan", "lion-air-jt610-jatuh1", "Lion Air JT610 jatuh", "lorem", "lorem", "published", time.Now(), time.Now(), time.Now()).
		AddRow(3, "fulan", "lion-air-jt610-jatuh2", "Lion Air JT610 jatuh", "lorem", "lorem", "published", time.Now(), time.Now(), time.Now()).
		AddRow(4, "fulan", "lion-air-jt610-jatuh3", "Lion Air JT610 jatuh", "lorem", "lorem", "draft", time.Now(), time.Now(), time.Now()).
		AddRow(5, "fulan", "lion-air-jt610-jatuh4", "Lion Air JT610 jatuh", "lorem", "lorem", "deleted", time.Now(), time.Now(), time.Now())

	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

	status := models.NewsStatusPublished
	results, err := repo.FetchByStatus(context.Background(), status)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), results[1].ID)
	assert.Equal(t, int64(3), results[2].ID)
}

func TestFetchNewsByParams(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := newsrepository.NewNewsSQLRepository(sqlxMockDb)
	query := "SELECT n.id, n.author, n.slug, n.title, n.description, n.content, n.status, n.published_at, n.created_at, n.updated_at FROM news n"
	rows := sqlmock.NewRows([]string{"id", "author", "slug", "title", "description", "content", "status", "published_at", "created_at", "updated_at"}).
		AddRow(1, "fulan", "kpk-bubar", "KPK Bubar", "lorem", "lorem", "draft", time.Now(), time.Now(), time.Now()).
		AddRow(2, "fulan", "lion-air-jt610-jatuh1", "Lion Air JT610 jatuh", "lorem", "lorem", "published", time.Now(), time.Now(), time.Now()).
		AddRow(3, "fulan", "lion-air-jt610-jatuh2", "Lion Air JT610 jatuh", "lorem", "lorem", "published", time.Now(), time.Now(), time.Now()).
		AddRow(4, "fulan", "lion-air-jt610-jatuh3", "Lion Air JT610 jatuh", "lorem", "lorem", "draft", time.Now(), time.Now(), time.Now())

	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

	status := models.NewsStatusPublished
	params := &models.FetchNewsParam{
		Pagination: &models.Pagination{Limit: 10, NextCursor: 0},
		Status:     status,
		TopicIDs:   []int64{1, 3},
	}
	results, err := repo.FetchByParams(context.Background(), params)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), results[0].ID)
	assert.Equal(t, int64(2), results[1].ID)
	assert.Equal(t, int64(3), results[2].ID)
}

func TestStoreNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := newsrepository.NewNewsSQLRepository(sqlxMockDb)
	query := "INSERT INTO \\`news\\`"
	lastId := int64(3)
	rowsAffected := int64(1)

	news := &models.News{
		Title:    "lorem",
		TopicIDs: []int64{1, 3},
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WillReturnResult(sqlmock.NewResult(lastId, rowsAffected))
	mock.ExpectPrepare("INSERT INTO news_topic \\(news_id,topic_id\\) VALUES \\(\\?,\\?\\),\\(\\?,\\?\\)").ExpectExec().WillReturnResult(sqlmock.NewResult(lastId, rowsAffected))
	mock.ExpectCommit()

	result, err := repo.Store(context.Background(), news)
	assert.NoError(t, err)
	assert.Equal(t, lastId, result)
}

func TestUpdateNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := newsrepository.NewNewsSQLRepository(sqlxMockDb)
	query := "UPDATE news SET title = \\?, updated_at = \\? WHERE id = \\?"
	lastId := int64(3)
	rowsAffected := int64(1)

	news := &models.News{
		ID:       3,
		Title:    "KPK Bubar",
		TopicIDs: []int64{},
	}

	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "author", "slug", "title", "description", "content", "status", "published_at", "created_at", "updated_at"}).
		AddRow(1, "fulan", "kpk-bubar", "KPK Bubar", "lorem", "lorem", "draft", now, now, now)

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WillReturnResult(sqlmock.NewResult(lastId, rowsAffected))
	mock.ExpectCommit()
	mock.ExpectPrepare("SELECT \\* FROM \\`news\\`").ExpectQuery().WillReturnRows(rows)

	result, err := repo.Update(context.Background(), news)
	assert.NoError(t, err)
	assert.Equal(t, news.Title, result.Title)
}

func TestDeleteNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxMockDb := sqlx.NewDb(db, "sqlmock")
	repo := newsrepository.NewNewsSQLRepository(sqlxMockDb)

	query := "DELETE FROM \\`news\\`"
	lastId := int64(1)
	rowsAffected := int64(1)

	mock.ExpectPrepare(query).ExpectExec().WillReturnResult(sqlmock.NewResult(lastId, rowsAffected))

	deleted, err := repo.Delete(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, true, deleted)
}
