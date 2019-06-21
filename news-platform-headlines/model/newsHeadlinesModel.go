package model

import "time"

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

type NewsheadlinesResponse struct {
	County       string    `json:"County"`
	Content      string    `json:"Content"`
	Published_At time.Time `json:"publishedAt"`
	Author       string    `json:"Author"`
}
