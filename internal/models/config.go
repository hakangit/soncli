package models

import "time"

// Config represents the application configuration
type Config struct {
	Sonarr  SonarrConfig  `mapstructure:"sonarr" yaml:"sonarr"`
	Sabnzbd SabnzbdConfig `mapstructure:"sabnzbd" yaml:"sabnzbd"`
	UI      UIConfig      `mapstructure:"ui" yaml:"ui"`
}

// SonarrConfig holds Sonarr connection settings
type SonarrConfig struct {
	Host    string        `mapstructure:"host" yaml:"host"`
	Port    int           `mapstructure:"port" yaml:"port"`
	APIKey  string        `mapstructure:"api_key" yaml:"api_key"`
	Timeout time.Duration `mapstructure:"timeout" yaml:"timeout"`
}

// SabnzbdConfig holds Sabnzbd connection settings
type SabnzbdConfig struct {
	Host     string        `mapstructure:"host" yaml:"host"`
	Port     int           `mapstructure:"port" yaml:"port"`
	APIKey   string        `mapstructure:"api_key" yaml:"api_key"`
	Username string        `mapstructure:"username" yaml:"username"`
	Password string        `mapstructure:"password" yaml:"password"`
	Timeout  time.Duration `mapstructure:"timeout" yaml:"timeout"`
}

// UIConfig holds UI-related settings
type UIConfig struct {
	Colors     bool `mapstructure:"colors" yaml:"colors"`
	MaxResults int  `mapstructure:"max_results" yaml:"max_results"`
}
