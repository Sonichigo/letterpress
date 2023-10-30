package main

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/sonichigo/letterpress/db"
	"github.com/sonichigo/letterpress/handler"
)

func main() {
	var err error
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dbConfig := db.Config{
		Host:     "postgres",
		Port:     5432,
		Username: "letterpress",
		Password: "letterpress_secrets",
		DbName:   "letterpress_db",
		Logger:   logger,
	}
	logger.Info().Interface("config", &dbConfig).Msg("config:")
	fmt.Println(dbConfig)
	dbInstance, err := db.Init(dbConfig)
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}
	logger.Info().Msg("Database connection established")

	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}

	h := handler.New(dbInstance, esClient, logger)
	router := gin.Default()
	rg := router.Group("/v2")
	h.Register(rg)
	router.Run(":8080")
}
