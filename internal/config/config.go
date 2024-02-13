package config

import "flag"

const (
	defaultDatabaseURI   = "postgresql://user:password@host:5432/gophkeeper"
	defaultServerAddress = "localhost:3200"
)

// Configuration хранит настройки сервера приложения.
type Configuration struct {
	ServerAddress string `env:"RUN_ADDRESS"`  // IP-адрес и порт работы gRPC-сервера приложения.
	DatabaseURI   string `env:"DATABASE_URI"` // Строка для подключения к базе данных приложения.
}

// NewConfiguration устанавливает настройки приложения.
func NewConfiguration() *Configuration {
	cfg := &Configuration{}

	flag.StringVar(&cfg.ServerAddress, "a", defaultServerAddress, "string with server address")
	flag.StringVar(&cfg.DatabaseURI, "d", defaultDatabaseURI, "string with database URI")
	flag.Parse()

	return cfg
}
