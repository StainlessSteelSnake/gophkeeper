package config

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cfgFileName      = "gophkeeper.yaml"
	cfgKeyphrase     = "keyphrase"
	cfgToken         = "token"
	cfgVersion       = "version"
	cfgBuildTime     = "build_time"
	cfgServerAddress = "server_address"
)

type config struct {
	v *viper.Viper
}

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
}

func ReadConfig() (Configurator, error) {
	result := &config{
		v: viper.New(),
	}

	result.v.SetDefault(cfgServerAddress, ":3200")

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

func (c *config) SetToken(token string) error {
	c.v.Set(cfgToken, token)
	return c.v.WriteConfigAs(cfgFileName)
}

func (c *config) GetToken() string {
	return c.v.GetString(cfgToken)
}

func (c *config) SetKeyPhrase(keyphrase string) error {
	c.v.Set(cfgKeyphrase, keyphrase)
	return c.v.WriteConfigAs(cfgFileName)
}

func (c *config) GetKeyPhrase() string {
	return c.v.GetString(cfgKeyphrase)
}

func (c *config) SetVersion(version, buildTime string) error {
	c.v.Set(cfgVersion, version)
	c.v.Set(cfgBuildTime, buildTime)
	return c.v.WriteConfigAs(cfgFileName)
}

func (c *config) GetVersion() (version, buildTime string) {
	version = c.v.GetString(cfgVersion)
	buildTime = c.v.GetString(cfgBuildTime)
	return version, buildTime
}

func (c *config) SetServerAddress(a string) error {
	c.v.Set(cfgServerAddress, a)
	return c.v.WriteConfigAs(cfgFileName)
}

func (c *config) GetServerAddress() string {
	return c.v.GetString(cfgServerAddress)
}

func (c *config) BindPFlag(key string, flag *pflag.Flag) error {
	return c.v.BindPFlag(key, flag)
}
