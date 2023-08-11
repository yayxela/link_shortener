package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
	)
}

func New() (config Config, err error) {
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
