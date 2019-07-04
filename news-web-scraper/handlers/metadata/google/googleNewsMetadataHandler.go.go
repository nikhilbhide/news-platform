package handler

import (
	"encoding/json"
	"github.com/nik/news-platform/news-web-scraper/repository"
	"net/http"
)

func GetCountryMetadata(w http.ResponseWriter, r *http.Request) {
	metadata := repository.GetCountiresMetadata()
	json.NewEncoder(w).Encode(metadata)
}
