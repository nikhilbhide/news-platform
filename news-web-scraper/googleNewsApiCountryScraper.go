package main

import "github.com/nik/news-platform/news-web-scraper/handlers/scraper"

//a set of countries scraped from web page
var countryMetadata map[string]string
var countires []string

//return a slice of countries
func extractCountries() []string {
	return countires
}

func main() {
	scraper.ScrapeAndStoreMetaData()
}
