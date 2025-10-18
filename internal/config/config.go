package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"sonarr-sabnzbd-cli/internal/models"
)

// LoadConfig loads the configuration from file and environment variables
func LoadConfig() (*models.Config, error) {
	// Set default values
	viper.SetDefault("sonarr.host", "localhost")
	viper.SetDefault("sonarr.port", 8989)
	viper.SetDefault("sonarr.api_key", "")
	viper.SetDefault("sonarr.timeout", 30*time.Second)
	viper.SetDefault("sabnzbd.host", "localhost")
	viper.SetDefault("sabnzbd.port", 8080)
	viper.SetDefault("sabnzbd.api_key", "")
	viper.SetDefault("sabnzbd.username", "")
	viper.SetDefault("sabnzbd.password", "")
	viper.SetDefault("sabnzbd.timeout", 30*time.Second)
	viper.SetDefault("ui.colors", true)
	viper.SetDefault("ui.max_results", 10)

	// Set config file name and paths
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Add config paths
	configDir, err := getConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directory: %w", err)
	}

	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	// Enable environment variable override
	viper.SetEnvPrefix("SONARR_CLI")
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, create default
			return createDefaultConfig(configDir)
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Check if we have the correct values, if not, recreate config
	if config.Sonarr.Host == "localhost" || config.Sonarr.APIKey == "" {
		return createDefaultConfig(configDir)
	}

	return &config, nil
}

// SaveConfig saves the configuration to file
func SaveConfig(config *models.Config) error {
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "config.yaml")

	viper.Set("sonarr", config.Sonarr)
	viper.Set("sabnzbd", config.Sabnzbd)
	viper.Set("ui", config.UI)

	return viper.WriteConfigAs(configPath)
}

// getConfigDir returns the configuration directory path
func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "sonarr-sabnzbd-cli")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// createDefaultConfig creates a default configuration file
func createDefaultConfig(_ string) (*models.Config, error) {
	// Create config with placeholder values - user must configure
	config := &models.Config{
		Sonarr: models.SonarrConfig{
			Host:    "localhost",
			Port:    8989,
			APIKey:  "", // User must set this
			Timeout: 30 * time.Second,
		},
		Sabnzbd: models.SabnzbdConfig{
			Host:     "localhost",
			Port:     8080,
			APIKey:   "", // User must set this
			Username: "",
			Password: "",
			Timeout:  30 * time.Second,
		},
		UI: models.UIConfig{
			Colors:     true,
			MaxResults: 10,
		},
	}

	// Save the default config
	if err := SaveConfig(config); err != nil {
		return nil, fmt.Errorf("failed to save default config: %w", err)
	}

	return config, nil
}
