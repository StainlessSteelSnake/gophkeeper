package config

import "flag"

const (
	defaultDatabaseURI   = "postgresql://user:password@host:5432/gophkeeper"
	defaultServerAddress = "localhost:3200"
)

type Configuration struct {
	ServerAddress string `env:"RUN_ADDRESS"`
	DatabaseURI   string `env:"DATABASE_URI"`
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{}

	flag.StringVar(&cfg.ServerAddress, "a", defaultServerAddress, "string with server address")
	flag.StringVar(&cfg.DatabaseURI, "d", defaultDatabaseURI, "string with database URI")
	flag.Parse()

	return cfg
}
