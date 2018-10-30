package main

import (
	"fmt"

	"gitlab.com/99ridho/news-api/domain/news/repository"
	"gitlab.com/99ridho/news-api/domain/news/usecase"

	"gitlab.com/99ridho/news-api/domain/news"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/99ridho/news-api/domain/news/delivery/http"
	"gitlab.com/99ridho/news-api/domain/topic/delivery/http"

	"gitlab.com/99ridho/news-api/domain/topic/usecase"

	"gitlab.com/99ridho/news-api/domain/topic/repository"

	"gitlab.com/99ridho/news-api/domain/topic"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gitlab.com/99ridho/news-api/app/database"
)

var conn *sqlx.DB

var topicRepo topic.TopicRepository
var topicUseCase topic.TopicUseCase
var newsRepo news.NewsRepository
var newsUseCase news.NewsUseCase

func init() {
	loadConfigurationFile()
	conn = appdatabase.InitializeDatabase()

	topicRepo = topicrepository.NewTopicSQLRepository(conn)
	topicUseCase = topicusecase.NewTopicUseCaseImplementation(topicRepo)
	newsRepo = newsrepository.NewNewsSQLRepository(conn)
	newsUseCase = newsusecase.NewNewsUseCaseImplementation(newsRepo)
}

func loadConfigurationFile() {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	defer conn.Close()

	r := echo.New()
	r.Use(middleware.Logger(), middleware.Recover())

	topichttpdelivery.InitializeTopicHandler(r, topicUseCase)
	newshttpdelivery.InitializeNewsHandler(r, newsUseCase)

	r.Start(viper.GetString("server.address"))
}
