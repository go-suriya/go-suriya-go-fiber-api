package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Name string `mapstructure:"NAME"`
	Env  string `mapstructure:"ENV"`
	Port int    `mapstructure:"PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	log.Printf("Current environment: %s", env)

	viper.SetConfigName(fmt.Sprintf(".env.%s", env))
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return config, nil
}

func (c *Config) GetPort() string {
	return fmt.Sprintf("0.0.0.0:%s", strconv.Itoa(c.Port))
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBPort,
		"disable",
		"public",
	)
}
