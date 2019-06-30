package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/nik/news-platform/common-platform/logger"
	"github.com/nik/news-platform/news-web-scraper/model"
	"github.com/nik/news-platform/news-web-scraper/repository"
	"io"
	"net/http"
	"os"
)

func ScrapeAndStoreMetaData(config model.Config) {
	countryMetadataCollection := scrapeWebsite(config)
	storeMetaData(countryMetadataCollection)
}

//scrapes the google website and extracts the country specific metadata
//the country metadata is stored in the map of country name and short name
func scrapeWebsite(config model.Config) []model.GoogleNewslinesMetadata {
	logger := logs.InitLogger(config.Logger.LogPath)
	var countryMetadata []model.GoogleNewslinesMetadata

	// Make HTTP request
	response, err := http.Get(config.WebsiteScraper.GoogleNewsMetadataAPI)
	logger.Info("Retreived the web page from google")

	//check for error
	if err != nil {
		logger.Fatal(err)
	}

	//close the resource
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		logger.Fatal("Error loading HTTP response body. ", err)
	}

	// Copy data from the response to standard output
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		logger.Fatal(err)
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

		//extract country short name
		countryShortName = element.Find("kbd").Text()

		if exists {
			//if shortname is retrieved from the current selection
			metaDataInstance := model.GoogleNewslinesMetadata{
				Country:   country,
				Shortname: countryShortName,
			}

			//add google news headlines instance to the collection
			countryMetadata = append(countryMetadata, metaDataInstance)
		}
	}

	//extract the countries from web page
	document.Find(".countries-and-categories .sources-container .source").Each(getCountries)
	logger.Info("Retreived and stored the metadata from the webpage")

	return countryMetadata
}

//stores the data into cassandra
func storeMetaData(inputMetadataCollection []model.GoogleNewslinesMetadata) {
	if inputMetadataCollection != nil && len(inputMetadataCollection) > 0 {
		repository.InsertGoogleNewsHeadlinesMetadata(inputMetadataCollection)
	}
}
