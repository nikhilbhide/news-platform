package services

import (
	"fmt"
	"github.com/nik/news-platform/common-platform/kafka"
	"time"
	"encoding/json"
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

func ListenAndProcessArticles() {
	//create a consumer
	consumer := kafka.CreateConsumer("article-scraper")

	//read message from the topic
	for {
		msg, _ := consumer.ReadMessage(-1)

		if(msg!=nil) {
			//transform the message into articles
			res := Newsheadlines{}
			if err := json.Unmarshal(msg.Value, &res); err != nil {
				panic(err)
			}

			for _, article := range res.Articles {
				fmt.Println(article)
			}
		}
	}
}