package services

import (
	"encoding/json"
	"github.com/nik/news-platform/common-platform/cassandra"
	"github.com/nik/news-platform/common-platform/kafka"
	"github.com/nik/news-platform/news-platform-headlines/model"
	"github.com/nik/news-platform/news-platform-headlines/repository"
	"time"
)

func ListenAndProcessArticles() {
	//create a consumer
	consumer := kafka.CreateConsumer("article-scraper")

	//read message from the topic
	for {
		msg, _ := consumer.ReadMessage(-1)

		if msg != nil {
			//extract the key as its the country name
			country := msg.Key
			//process messages
			if msg != nil {
				//transform the message into articles
				res := model.Newsheadlines{}
				if err := json.Unmarshal(msg.Value, &res); err != nil {
					panic(err)
				}

				//connect to cassandra cluster and insert the articles into news_headlines
				session := cassandra.ConnectToCluster()
				defer cassandra.CloseSession(session)
				repository.Insert(session, string(country), res)
			}

			//sleep for 1000
			time.Sleep(time.Microsecond)
		}
	}
}
