package router

import (
	"github.com/gorilla/mux"
	"github.com/nik/news-platform/news-platform-headlines/services/handler"
)

func CreateRouter() *mux.Router {
	var (
		router    = mux.NewRouter()
		apiRouter = router.PathPrefix("/news-platform/").Subrouter()
	)

	apiRouter.HandleFunc("/headlines", handler.GetTodaysHeadlines).Methods("GET")
	return router
}
