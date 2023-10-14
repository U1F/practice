// config/config.go
package config

import (
	"encoding/json"
	"os"
)

// Config represents the application configuration.
type Config struct {
	ServerAddress    string   `json:"server_address"`
	AllowedOrigins   []string `json:"allowed_origins"`
	AllowCredentials bool     `json:"allow_credentials"`
	AllowedMethods   []string `json:"allowed_methods"`
}

// LoadConfig reads the configuration from a JSON file.
func LoadConfig(filePath string) (Config, error) {
	var cfg Config

	file, err := os.Open(filePath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
