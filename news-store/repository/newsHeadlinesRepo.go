package repository

import (
	"github.com/gocql/gocql"
	"github.com/nik/news-platform/news-store/model"
)

const headlinesTable = "news_headlines"
const insertQuery = "INSERT INTO news_headlines (country, author, content, date, description, url, published_at,title, source_id, source_name) VALUES (?, ?, ?,?, ?, ?, ?,?,?)"

//method to insert the records into the table
func Insert(session *gocql.Session, country string, input model.Newsheadlines) {
	for _, article := range input.Articles {
		if err := session.Query(insertQuery,
			country, article.Author, article.Content, article.PublishedAt.Format("01-02-2006"), article.Description, article.URL, article.PublishedAt, article.Title, article.Source.ID, article.Source.Name).Exec(); err != nil {
			panic(err)
		}
	}
}
