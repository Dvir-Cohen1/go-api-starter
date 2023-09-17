package config

import (
	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	Port int `mapstructure:"PORT"`
	// Add other configuration parameters as needed
}

// LoadConfig loads the application configuration from environment variables or config files.
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config") // Name of your config file (e.g., config.yaml)
	viper.AddConfigPath(".")      // Search in the current directory for the config file
	viper.AutomaticEnv()          // Read environment variables

	var config AppConfig
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
