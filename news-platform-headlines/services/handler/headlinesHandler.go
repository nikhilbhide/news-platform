package handler

import (
	"encoding/json"
	"fmt"
	"github.com/nik/news-platform/common-platform/cassandra"
	"github.com/nik/news-platform/news-platform-headlines/repository"
	"net/http"
	"time"
)

func GetTodaysHeadlines(w http.ResponseWriter, r *http.Request) {
	//extract query parameters
	queryMap := r.URL.Query()

	//check whether query parameter query is present
	country := queryMap.Get("country")
	fmt.Println(country)

	//connect to cassandra cluster and insert the articles into news_headlines
	session := cassandra.ConnectToCluster()
	defer cassandra.CloseSession(session)
	date := time.Now()
	//Format MM-DD-YYYY
	fmt.Println(date.Format("01-02-2006"))
	headlines := repository.GetHeadlinesByCountry(session, country)
	json.NewEncoder(w).Encode(headlines)
}
