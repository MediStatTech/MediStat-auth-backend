package config

import "github.com/spf13/viper"

type Config struct {
	JWTTokenSecret string `mapstructure:"JWT_TOKEN_SECRET" validate:"required"`
	HTTPPort       string `mapstructure:"HTTP_PORT" validate:"required"`
	PostgresDSN    string `mapstructure:"POSTGRES_DSN" validate:"required"`
	JWTSecret      string `mapstructure:"JWT_SECRET" validate:"required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
