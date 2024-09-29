package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotToken string
}

func LoadConfig(path string) (*Config, error) {

	config := new(Config)

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func MustLoadConfig(path string) *Config {
	config, err := LoadConfig(path)
	if err != nil {
		panic(err)
	}

	return config
}
