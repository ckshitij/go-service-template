package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the overall application configuration
type Config struct {
	Server    ServerConfig              `mapstructure:"server"`
	Databases map[string]DatabaseConfig `mapstructure:"databases"`
	Redis     RedisConfig               `mapstructure:"redis"`
}

// ServerConfig holds server-related configurations
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// DatabaseConfig holds individual database configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// RedisConfig holds Redis-related configurations
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// LoadConfig loads the configuration file using Viper
func LoadConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)

	// Enable environment variable overrides
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return &config, nil
}
