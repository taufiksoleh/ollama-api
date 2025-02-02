package config

import (
	"os"
)

type Config struct {
	ServerAddress  string
	OllamaEndpoint string
}

func Load() (*Config, error) {
	serverAddr := os.Getenv("SERVER_ADDRESS")
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	ollamaEndpoint := os.Getenv("OLLAMA_ENDPOINT")
	if ollamaEndpoint == "" {
		ollamaEndpoint = "http://localhost:11434"
	}

	return &Config{
		ServerAddress:  serverAddr,
		OllamaEndpoint: ollamaEndpoint,
	}, nil
}
