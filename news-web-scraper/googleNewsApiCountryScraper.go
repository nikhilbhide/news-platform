package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
)

//a set of countries scraped from web page
var countrySet map[string]bool
var countires []string

//return a slice of countries
func GetCountries() []string {
	return countires
}

func main() {
	countrySet = make(map[string]bool) // New empty set

	// Make HTTP request
	response, err := http.Get("https://newsapi.org/sources")
	//response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Make request
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find all links and process them with the function
	getCountries := func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		country, exists := element.Attr("data-country")
		if exists {
			countrySet[country] = true
		}
	}

	//extract the countries from web page
	document.Find("div").Each(getCountries)

	//iterate over map and retrieve countires
	for key, _ := range countrySet {
		countires = append(countires, key)
	}
}
