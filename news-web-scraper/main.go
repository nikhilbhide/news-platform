package main

import (
	"github.com/nik/news-platform/common-platform/logger"
	"github.com/nik/news-platform/news-web-scraper/config"
	"github.com/nik/news-platform/news-web-scraper/handlers/metadata/google"
	"github.com/nik/news-platform/news-web-scraper/handlers/scraper"
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
	/*middlewareManager := alice.New(
	transactionId.AddIDRequestMiddleware,transactionId.AddIDRequestMiddleware).
	Then(router.CreateRouter())
	*/
	logger.Info("http handlers are initialized")

	//start listening
	http.HandleFunc("/view/", handler.GetCountryMetadata)

	http.ListenAndServe("localhost:8083", nil)
}
