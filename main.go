package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func init() {
	loadConfigurationFile()
	conn = appdatabase.InitializeDatabase()

	topicRepo = topicrepository.NewTopicSQLRepository(conn)
	topicUseCase = topicusecase.NewTopicUseCaseImplementation(topicRepo)
}

func loadConfigurationFile() {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	r := echo.New()
	r.Use(middleware.Logger(), middleware.Recover())

	topichttpdelivery.InitializeTopicHandler(r, topicUseCase)

	r.Start(viper.GetString("server.address"))
}
