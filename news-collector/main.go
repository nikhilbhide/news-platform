package main

import (
	"github.com/nik/news-platform/news-collector/config"
	"github.com/nik/news-platform/news-collector/services"
)

func main() {
	config := config.LoadConfiguration("news-collector/config/config.json")
	services.Init(config.GoogleNewsAPI.Url, config.GoogleNewsAPI.APIKey)
	services.ScrapeNewsHeadlines(config)
}
