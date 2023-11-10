package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	cfgFileName  = "gophkeeper.yaml"
	cfgKeyphrase = "keyphrase"
	cfgToken     = "token"
	cfgVersion   = "version"
	cfgBuildTime = "build_time"
)

type config struct {
	v *viper.Viper
}

type Configurator interface {
	SetToken(string) error
	GetToken() string
	SetKeyPhrase(string) error
	GetKeyPhrase() string
	SetVersion(string, string)
	GetVersion() (string, string)
}

func init() {
	viper.SetDefault(cfgToken, "")
	viper.SetDefault(cfgKeyphrase, "")
}

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
		} else {
			return nil, err
		}
	}

	return result, nil
}

func (c *config) SetToken(token string) error {
	c.v.Set(cfgToken, token)
	err := c.v.WriteConfigAs(cfgFileName)
	return err
}

func (c *config) GetToken() string {
	return c.v.GetString(cfgToken)
}

func (c *config) SetKeyPhrase(keyphrase string) error {
	c.v.Set(cfgKeyphrase, keyphrase)
	err := c.v.WriteConfigAs(cfgFileName)
	return err
}

func (c *config) GetKeyPhrase() string {
	return c.v.GetString(cfgKeyphrase)
}

func (c *config) SetVersion(version, buildTime string) {
	c.v.Set(cfgVersion, version)
	c.v.Set(cfgBuildTime, buildTime)
}

func (c *config) GetVersion() (version, buildTime string) {
	version = c.v.GetString(cfgVersion)
	buildTime = c.v.GetString(cfgBuildTime)
	return version, buildTime
}
