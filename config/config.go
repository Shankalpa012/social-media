package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func New() Config {

	env := Config{}
	//setting the congig file
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		os.Exit(1)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		os.Exit(1)
	}

	return env
}
