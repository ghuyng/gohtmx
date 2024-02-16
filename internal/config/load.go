package config

import (
	"github.com/spf13/viper"
)

func Load(fileName string) (*Config, error) {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
