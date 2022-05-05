package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	AppName string `yaml:"appName"`
	Port    string `yaml:"port"`
	Mode    string `yaml:"mode"`
	Version string `yaml:"version"`
}

type DatabaseConfiguration struct {
	Driver   string `yaml:"driver"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("error reading config file")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot unmarshal into struct")
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
