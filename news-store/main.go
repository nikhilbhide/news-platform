package main

import (
	"github.com/nik/news-platform/news-store/router"
	"github.com/nik/news-platform/news-store/services"
)

func main() {
	go services.ListenAndProcessArticles()
	router.CreateRouter()
}
