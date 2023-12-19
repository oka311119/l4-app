package config

import (
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
