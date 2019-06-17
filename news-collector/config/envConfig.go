package config

import (
	"github.com/nik/news-platform/news-collector/model"
	"os"
)
import "fmt"
import "encoding/json"

//Load configuration from the file
func LoadConfiguration(file string) model.Config {
	var config model.Config

	//open file and read configuration
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
