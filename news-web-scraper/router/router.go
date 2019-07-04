package router

import (
	"github.com/gorilla/mux"
	"github.com/nik/news-platform/news-platform-headlines/services/handler"
)

func CreateRouter() *mux.Router {
	var (
		router            = mux.NewRouter()
		apiRouterMetadata = router.PathPrefix("/news-web-scraper/metadata").Subrouter()
	)

	apiRouterMetadata.HandleFunc("/googlenews/countires", handler.GetTodaysHeadlines).Methods("GET")
	return router
}
