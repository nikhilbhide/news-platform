package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
)

func ScrapeAndStoreMetaData() {

}

//scrapes the google website and extracts the country specific metadata
//the country metadata is stored in the map of country name and short name
func scrapeWebsite() {
	var countryMetadata map[string]string

	// Make HTTP request
	response, err := http.Get("https://newsapi.org/sources")

	//check for error
	if err != nil {
		log.Fatal(err)
	}

	//close the resource
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Copy data from the response to standard output
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find country and short country name
	getCountries := func(index int, element *goquery.Selection) {
		var (
			country          string
			countryShortName string
			exists           bool
		)
		element.Find(".name").Each(func(index int, element *goquery.Selection) {
			country, exists = element.Attr("title")
		})

		countryShortName = element.Find("kbd").Text()

		if exists {
			countryMetadata[country] = countryShortName
		}
	}

	//extract the countries from web page
	document.Find(".countries-and-categories .sources-container .source").Each(getCountries)
}

func storeMetaData(map[string]string) {

}
