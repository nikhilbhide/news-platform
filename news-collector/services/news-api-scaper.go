package services

import (
	"encoding/json"
	"fmt"
	"github.com/nik/news-platform/common-platform/kafka"
	"github.com/nik/news-platform/news-collector/model"
	"io/ioutil"
	"net/http"
)

var NewsAPIUrl string
var ApiKey string

func Init(url string, apiKey string) {
	NewsAPIUrl = url
	ApiKey = apiKey
}

//scrapes news headlines
func ScrapeNewsHeadlines(config model.Config) {
	fmt.Println("Starting the application...")

	countries := []string{"us", "in", "gb"}

	for _, country := range countries {
		newsAPIUrl := fmt.Sprintf(NewsAPIUrl, country, ApiKey)

		//get response by countries
		response, err := http.Get(newsAPIUrl)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
			res := model.Newsheadlines{}

			//convert the data into json?
			//Is it required? We will revisit and address the issue
			if err := json.Unmarshal(data, &res); err != nil {
				panic(err)
			}

			//send data to the topic
			kafka.SendMessage(config.Topic, country, data)
		}
	}
}
