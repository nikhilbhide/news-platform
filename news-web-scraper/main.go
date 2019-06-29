package main

import "github.com/nik/news-platform/news-web-scraper/handlers/scraper"

func main() {
	//start scraping google website and store metadata
	scraper.ScrapeAndStoreMetaData()
}
