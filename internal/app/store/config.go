package store

type Config struct {
	Datanase_url string `json: database_url`
}

func NewConfig() *Config {
	return &Config
}
