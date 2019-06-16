package services

import (
	"encoding/json"
	"fmt"
	"github.com/nik/news-platform/common-platform/kafka"
	"io/ioutil"
	"net/http"
	"time"
)

type Newsheadlines struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}

var newsAPIUrl string
var apiKey string

func Init(url string, apiKey string) {
	newsAPIUrl = url
	apiKey = apiKey
}

func ScrapeNewsHeadlines() {
	fmt.Println("Starting the application...")

	countries := []string{"us", "in", "gb"}

	for _, country := range countries {
		newsAPIUrl := fmt.Sprintf(newsAPIUrl, country, apiKey)

		//get response by countries
		response, err := http.Get(newsAPIUrl)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
			res := Newsheadlines{}

			//convert the data into json?
			//Is it required? We will revisit and address the issue
			if err := json.Unmarshal(data, &res); err != nil {
				panic(err)
			}

			//send data to the topic
			kafka.SendMessage("article-scraper", country, data)
		}
	}
}
