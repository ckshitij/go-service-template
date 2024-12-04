package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Config represents the overall application configuration
type Config struct {
	Server    ServerConfig
	Databases map[string]DatabaseConfig
	Redis     RedisConfig
}

// ServerConfig holds server-related configurations
type ServerConfig struct {
	Host string
	Port int
}

// DatabaseConfig holds individual database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// RedisConfig holds Redis-related configurations
type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load the .env file

	// Enable environment variable overrides
	viper.AutomaticEnv()

	// Map environment variables to struct fields
	config := &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnvAsInt("SERVER_PORT", 8080),
		},
		Databases: map[string]DatabaseConfig{
			"users": {
				Host:     getEnv("DATABASE_USERS_HOST", "localhost"),
				Port:     getEnvAsInt("DATABASE_USERS_PORT", 5432),
				User:     viper.GetString("DATABASE_USERS_USER"),
				Password: viper.GetString("DATABASE_USERS_PASSWORD"),
				Database: viper.GetString("DATABASE_USERS_DATABASE"),
			},
		},
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: viper.GetString("REDIS_PASSWORD"),
		},
	}

	return config, nil
}

// Helper function to get environment variable as integer
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// Helper function to get environment variable as integer
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
