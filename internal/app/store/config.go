package store

type Config struct {
	Database_url string `json: database_url`
}

func NewConfig() *Config {
	return &Config{}
}
