package repository

import (
	"github.com/gocql/gocql"
	"github.com/nik/news-platform/common-platform/cassandra"
	"github.com/nik/news-platform/news-web-scraper/model"
	"log"
)

const headlinesTable = "news_headlines"
const insertQuery = "INSERT INTO news_headlines_web_country_metadata (country, shortname) VALUES (?, ?)"

//method to insert the records into the table
func InsertGoogleNewsHeadlinesMetadata(response []model.GoogleNewslinesMetadata) {
	//create a session object
	session := cassandra.ConnectToCluster()
	defer cassandra.CloseSession(session)

	//create a batch
	batch := gocql.NewBatch(gocql.LoggedBatch)
	for counter := 0; counter < len(response); counter++ {
		batch.Query(insertQuery, response[counter].Country, response[counter].Shortname)
	}

	//execute a batch of cassandra
	err := session.ExecuteBatch(batch)
	if err != nil {
		log.Fatal("Insert failed", err)
	}
}
