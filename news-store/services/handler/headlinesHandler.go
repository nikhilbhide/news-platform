package handler

import (
	"fmt"
	"net/http"
)

func HeadlinesHandler(w http.ResponseWriter, r *http.Request) {
	//extract query parameters
	queryMap := r.URL.Query()

	//check whether query parameter query is present
	country := queryMap.Get("country")
	fmt.Println(country)
}
