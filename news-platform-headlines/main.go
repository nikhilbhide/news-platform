package main

import (
	"fmt"
	"github.com/justinas/alice"
	"github.com/nik/news-platform/news-platform-headlines/config"
	"github.com/nik/news-platform/news-platform-headlines/interceptor"
	"github.com/nik/news-platform/news-platform-headlines/router"
	"github.com/nik/news-platform/news-platform-headlines/services"
	"net/http"
)

func main() {
	go services.ListenAndProcessArticles()

	//load the configuration
	config := config.LoadConfiguration("news-platform-headlines/config/config.json")

	// setting up web server middlewares
	middlewareManager := alice.New(
		transactionId.AddIDRequestMiddleware).
		Then(router.CreateRouter())

	//start listening
	err := http.ListenAndServe(config.ListernURL, middlewareManager)
	fmt.Println("Stopping the application : %v", err)
}
