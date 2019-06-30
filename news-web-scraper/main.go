package main

import (
	"github.com/nik/news-platform/common-platform/logger"
	"github.com/nik/news-platform/news-web-scraper/config"
	"github.com/nik/news-platform/news-web-scraper/handlers/scraper"
)

func main() {
	//load the configuration
	config := config.LoadConfiguration("news-web-scraper/config/config.json")
	//set the logger to use advanced logging
	logger := logs.InitLogger(config.Logger.LogPath)
	logger.Info("Bootstrapping the application")
	//start scraping google website and store metadata
	scraper.ScrapeAndStoreMetaData(config)
}
