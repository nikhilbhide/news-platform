package model

//configuration structure
type Config struct {
	Cassandra     Cassandra     `json:"Cassandra"`
	KafkaBrokers  string        `json:"KafkaBrokers"`
	Topic         string        `json:"Topic"`
	Version       Version       `json:"Version"`
	GoogleNewsAPI GoogleNewsAPI `json:"GoogleNewsAPI"`
	ListernURL    string        `json:"ListernURL"`
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

type Logger struct {
	LogPath string `json:"LogPath"`
}
