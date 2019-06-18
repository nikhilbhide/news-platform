package router

import (
	"github.com/gorilla/mux"
	"github.com/nik/news-platform/news-store/services/handler"
	"net/http"
	"time"
)

func CreateRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/news-platform/top-headlines", handler.HeadlinesHandler)
	srv := &http.Server{
		Addr: "0.0.0.0:8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	srv.ListenAndServe()
}
