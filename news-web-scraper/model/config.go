package model

//configuration structure
type Config struct {
	Cassandra struct {
		Host     string `json:"Host"`
		Keyspace string `json:Keyspace'`
	} `json:"Cassandra"`
	KafkaBrokers   string         `json:"KafkaBrokers"`
	Version        Version        `json:"Version"`
	WebsiteScraper WebsiteScraper `json:"WebsiteScraper"`
	Logger         Logger         `json:"Logger"`
}

type Version struct {
	SolutionName    string `json:"SolutionName"    envconfig:"TKS_SOLUTION_NAME"`
	ServiceName     string `json:"ServiceName"     envconfig:"TKS_SERVICE_NAME"`
	ServiceProvider string `json:"ServiceProvider" envconfig:"TKS_SERVICE_PROVIDER"`
}

type WebsiteScraper struct {
	GoogleNewsMetadataAPI string `json:"GoogleNewsMetadataAPI"    envconfig:"GOOGLE_NEWS_METADATA_API"`
}

type Logger struct {
	LogPath string `json:"LogPath"`
}
