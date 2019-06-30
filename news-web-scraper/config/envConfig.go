package config

import (
	"github.com/nik/news-platform/news-web-scraper/model"
	"os"
)
import "fmt"
import "encoding/json"

//load configuration from the config file
func LoadConfiguration(file string) model.Config {
	var config model.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
