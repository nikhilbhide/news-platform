package model

//configuration structure
type Config struct {
	Cassandra struct {
		Host     string `json:"Host"`
		Keyspace string `json:Keyspace'`
	} `json:"Cassandra"`
	KafkaBrokers string  `json:"Host"`
	Port         string  `json:"Port"`
	Version      Version `json:"Version"`
}

type Version struct {
	SolutionName    string `json:"SolutionName"    envconfig:"TKS_SOLUTION_NAME"`
	ServiceName     string `json:"ServiceName"     envconfig:"TKS_SERVICE_NAME"`
	ServiceProvider string `json:"ServiceProvider" envconfig:"TKS_SERVICE_PROVIDER"`
}
