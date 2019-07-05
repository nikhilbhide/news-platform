package main

import (
	"github.com/justinas/alice"
	"github.com/nik/news-platform/common-platform/logger"
	"github.com/nik/news-platform/news-web-scraper/config"
	"github.com/nik/news-platform/news-web-scraper/handlers/scraper"
	"github.com/nik/news-platform/news-web-scraper/interceptor"
	"github.com/nik/news-platform/news-web-scraper/router"
	"net/http"
)

func main() {
	//load the configuration
	config := config.LoadConfiguration("news-web-scraper/config/config.json")
	//set the logger to use advanced logging
	logger := logs.InitLogger(config.Logger.LogPath)
	logger.Info("Bootstrapping the application")
	//start scraping google website and store metadata
	scraper.ScrapeAndStoreMetaData(config)

	//setting up web server middlewares
	middlewareManager := alice.New(
		transactionId.AddIDRequestMiddleware).
		Then(router.CreateRouter())

	logger.Info("http handlers are initialized")
	error := http.ListenAndServe(config.ListernURL, middlewareManager)
	logger.Infof("Stopping the application %v", error)
}
