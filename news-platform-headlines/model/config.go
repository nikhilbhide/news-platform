package model

//configuration structure
type Config struct {
	Cassandra struct {
		Host     string `json:"Host"`
		Keyspace string `json:Keyspace'`
	} `json:"Cassandra"`
	KafkaBrokers string  `json:"Host"`
	ListernURL   string  `json:"ListernURL"`
	Version      Version `json:"Version"`
	Logger       Logger  `json:"Logger"`
}

type Logger struct {
	LogPath string `json:"LogPath"`
}

type Version struct {
	SolutionName    string `json:"SolutionName"    envconfig:"TKS_SOLUTION_NAME"`
	ServiceName     string `json:"ServiceName"     envconfig:"TKS_SERVICE_NAME"`
	ServiceProvider string `json:"ServiceProvider" envconfig:"TKS_SERVICE_PROVIDER"`
}
