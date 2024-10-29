package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl       string `mapstructure:"DB_URL"`
	HTTPPort    int    `mapstructure:"HTTP_PORT"`
	GatewayAUrl string `mapstructure:"GATEWAY_A_URL"`
	GatewayBUrl string `mapstructure:"GATEWAY_B_URL"`
	Timeout     int    `mapstructure:"TIMEOUT"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
