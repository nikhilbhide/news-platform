package config

import (
	"os"
)
import "fmt"
import "encoding/json"

//configuration structure
type Config struct {
	Cassandra     Cassandra     `json:"Cassandra"`
	KafkaBrokers  string        `json:"KafkaBrokers"`
	Version       Version       `json:"Version"`
	GoogleNewsAPI GoogleNewsAPI `json:"GoogleNewsAPI"`
}

//google news api configuration
type GoogleNewsAPI struct {
	Url    string `json:"Url" envconfig:"Url"`
	APIKey string `json:"APIKey" envconfig:"APIKey"`
}

//cassandra configuration
type Cassandra struct {
	Host     string `json:"Host" envconfig:CASSANDRA_HOST"`
	Keyspace string `json:"Keyspace" envconfig:"CASSANDRA_KEYSPACE"`
}

//microservice version structure
type Version struct {
	SolutionName    string `json:"SolutionName"    envconfig:"NC_SOLUTION_NAME"`
	ServiceName     string `json:"ServiceName"     envconfig:"NC_SERVICE_NAME"`
	ServiceProvider string `json:"ServiceProvider" envconfig:"NC_SERVICE_PROVIDER"`
	Number          string `json:"Number" envconfig:"NC_NUMBER"`
}

//Load configuration from the file
func LoadConfiguration(file string) Config {
	var config Config

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
