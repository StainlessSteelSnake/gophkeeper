package config

const (
	defaultDatabaseURI   = "postgresql://gophkeeper_admin:a1s9u8k3a@localhost:5432/gophkeeper"
	defaultServerAddress = "localhost:8081"
)

type Configuration struct {
	ServerAddress string `env:"RUN_ADDRESS"`
	DatabaseURI   string `env:"DATABASE_URI"`
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{}

	cfg.ServerAddress = defaultServerAddress
	cfg.DatabaseURI = defaultDatabaseURI

	return cfg
}
