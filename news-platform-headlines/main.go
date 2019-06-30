package main

import (
	"github.com/justinas/alice"
	"github.com/nik/news-platform/common-platform/logger"
	"github.com/nik/news-platform/news-platform-headlines/config"
	"github.com/nik/news-platform/news-platform-headlines/interceptor"
	"github.com/nik/news-platform/news-platform-headlines/router"
	"github.com/nik/news-platform/news-platform-headlines/services"
	"net/http"
)

func main() {
	//load the configuration
	config := config.LoadConfiguration("news-platform-headlines/config/config.json")

	//set the logger to use advanced logging
	logger := logs.InitLogger(config.Logger.LogPath)
	logger.Info("Bootstrapping the application")

	//create a kafka consumer
	go services.ListenAndProcessArticles()

	logger.Info("Kafka consumer started")

	logger.Info("Configuration is loaded")

	//setting up web server middlewares
	middlewareManager := alice.New(
		transactionId.AddIDRequestMiddleware).
		Then(router.CreateRouter())

	logger.Info("http handlers are initialized")

	//start listening
	error := http.ListenAndServe(config.ListernURL, middlewareManager)
	logger.Infof("Stopping the application %v", error)
}
