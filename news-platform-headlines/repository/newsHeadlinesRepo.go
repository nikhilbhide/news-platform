package repository

import (
	"github.com/gocql/gocql"
	"github.com/nik/news-platform/news-platform-headlines/model"
	"time"
)

const headlinesTable = "news_headlines"
const insertQuery = "INSERT INTO news_headlines (country, author, content, date, description, url, published_at,title, source_id, source_name) VALUES (?, ?, ?,?, ?, ?, ?,?,?,?)"
const selectQuery = "SELECT country,content,author,published_at from news_headlines where country = ?"

//method to insert the records into the table
func Insert(session *gocql.Session, country string, input model.Newsheadlines) {
	for _, article := range input.Articles {
		if err := session.Query(insertQuery,
			country, article.Author, article.Content, article.PublishedAt.Format("01-02-2006"), article.Description, article.URL, article.PublishedAt, article.Title, article.Source.ID, article.Source.Name).Exec(); err != nil {
			panic(err)
		}
	}
}

//gets the headlines from the database
func GetHeadlinesByCountry(session *gocql.Session, country string) []model.NewsheadlinesResponse {
	var (
		content      string
		published_at time.Time
		author       string
	)

	var headlines []model.NewsheadlinesResponse

	//gets the data from table
	iter := session.Query(selectQuery, country).Iter()

	//iterate over the results and populate headlines response
	for iter.Scan(&country, &content, &author, &published_at) {
		newsHeadlinesInstance := model.NewsheadlinesResponse{
			County:       country,
			Content:      content,
			Published_At: published_at,
			Author:       author,
		}

		//append the headlines
		headlines = append(headlines, newsHeadlinesInstance)
	}

	return headlines
}
