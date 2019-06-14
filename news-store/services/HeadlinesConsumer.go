package services

import (
	"encoding/json"
	"fmt"
	"github.com/nik/news-platform/common-platform/cassandra"
	"github.com/nik/news-platform/common-platform/kafka"
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

func ListenAndProcessArticles() {
	//create a consumer
	consumer := kafka.CreateConsumer("article-scraper")

	//read message from the topic
	for {
		msg, _ := consumer.ReadMessage(-1)

		//process messages
		if msg != nil {
			//transform the message into articles
			res := Newsheadlines{}
			if err := json.Unmarshal(msg.Value, &res); err != nil {
				panic(err)
			}

			//connect to cassandra cluster and insert the articles into news_headlines
			session := cassandra.ConnectToCluster()
			defer cassandra.CloseSession(session)
			for _, article := range res.Articles {
				if err := session.Query(`
      					INSERT INTO news_headlines (country, author, content, date, description, url, published_at) VALUES (?, ?, ?,?, ?, ?, ?)`,
					"us", article.Author, article.Content, article.PublishedAt.Format("01-02-2006"), article.Description, article.URL, article.PublishedAt).Exec(); err != nil {
					fmt.Println(err)
				} else {

				}
			}
		}

		//sleep for 1000
		time.Sleep(time.Microsecond)
	}
}
