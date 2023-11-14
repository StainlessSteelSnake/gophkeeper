package config

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cfgFileName      = "gophkeeper.yaml" // Название файла для хранения настроек клиентского приложения.
	cfgKeyphrase     = "keyphrase"       // Название настройки для хранения ключа шифрования.
	cfgToken         = "token"           // Название настройки для хранения токена авторизованного пользователя.
	cfgVersion       = "version"         // Название настройки для хранения версии клиентского приложения.
	cfgBuildTime     = "build_time"      // Название настройки для хранения даты и времени сборки клиентского приложения.
	cfgServerAddress = "server_address"  // Название настройки для хранения IP-адреса и порта сервера приложения.
)

// config хранит ссылку на контроллер управления настройками приложения.
type config struct {
	v *viper.Viper
}

// Configurator описывает методы получения и изменения настроек приложения.
type Configurator interface {
	SetToken(string) error
	GetToken() string
	SetKeyPhrase(string) error
	GetKeyPhrase() string
	SetVersion(string, string) error
	GetVersion() (string, string)
	SetServerAddress(string) error
	GetServerAddress() string
	BindPFlag(string, *pflag.Flag) error
}

func init() {
	viper.SetDefault(cfgToken, "")
	viper.SetDefault(cfgKeyphrase, "")
	viper.SetDefault(cfgServerAddress, ":3200")
}

// ReadConfig считывает настройки клиентского приложения из файла.
func ReadConfig() (Configurator, error) {
	result := &config{
		v: viper.New(),
	}

	result.v.SetConfigName("gophkeeper")
	result.v.SetConfigType("yaml")
	result.v.AddConfigPath(".")
	err := result.v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println(err)
			err = result.v.WriteConfigAs(cfgFileName)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return result, nil
}

// SetToken сохраняет в настройках приложения токен авторизованного пользователя.
func (c *config) SetToken(token string) error {
	c.v.Set(cfgToken, token)
	return c.v.WriteConfigAs(cfgFileName)
}

// GetToken получает из настроек приложения токен авторизованного пользователя.
func (c *config) GetToken() string {
	return c.v.GetString(cfgToken)
}

// SetKeyPhrase сохраняет в настройках приложения ключ шифрования данных.
func (c *config) SetKeyPhrase(keyphrase string) error {
	c.v.Set(cfgKeyphrase, keyphrase)
	return c.v.WriteConfigAs(cfgFileName)
}

// GetKeyPhrase получает из настроек приложения ключ шифрования данных.
func (c *config) GetKeyPhrase() string {
	return c.v.GetString(cfgKeyphrase)
}

// SetVersion сохраняет в настройках приложения его версию.
func (c *config) SetVersion(version, buildTime string) error {
	c.v.Set(cfgVersion, version)
	c.v.Set(cfgBuildTime, buildTime)
	return c.v.WriteConfigAs(cfgFileName)
}

// GetVersion получает из настройек приложения его версию.
func (c *config) GetVersion() (version, buildTime string) {
	version = c.v.GetString(cfgVersion)
	buildTime = c.v.GetString(cfgBuildTime)
	return version, buildTime
}

// SetServerAddress сохраняет в настройках приложения IP-адрес и порт для подключения к серверу приложения.
func (c *config) SetServerAddress(a string) error {
	c.v.Set(cfgServerAddress, a)
	return c.v.WriteConfigAs(cfgFileName)
}

// GetServerAddress получает из настроек приложения IP-адрес и порт для подключения к серверу приложения.
func (c *config) GetServerAddress() string {
	return c.v.GetString(cfgServerAddress)
}

// BindPFlag связывает флаг команды приложения и настройку приложения.
func (c *config) BindPFlag(key string, flag *pflag.Flag) error {
	return c.v.BindPFlag(key, flag)
}
