package store

type Store struct {
	config *Config
}

func New(config *Config) *Store {
	return &Config{
		config: config,
	}
}

func (s *Store) Open() {
	return nil
}

func (s *Store) Close() {
	return
}
