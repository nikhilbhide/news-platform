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

func ScrapeNewsHeadlines() {
	fmt.Println("Starting the application...")

	countries := []string{"us", "in", "gb"}
	url := "https://newsapi.org/v2/top-headlines?country=%s&apiKey=456f513b38d14b4a8c96fd267274c62d"

	for _, country := range countries {
		countryUrl := fmt.Sprintf(url, country)

		response, err := http.Get(countryUrl)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
			res := Newsheadlines{}

			if err := json.Unmarshal(data, &res); err != nil {
				panic(err)
			}

			//send data to the topic
			kafka.SendMessage("article-scraper", url, data)
		}
	}
}
