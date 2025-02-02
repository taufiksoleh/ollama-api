package main

import (
	"log"
	"net/http"

	"github.com/taufiksoleh/ollama-api/config"
	"github.com/taufiksoleh/ollama-api/handler"
	"github.com/taufiksoleh/ollama-api/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize service
	ollamaService := service.NewOllamaService(cfg.OllamaEndpoint)

	// Initialize handler
	ollamaHandler := handler.NewOllamaHandler(ollamaService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/generate", ollamaHandler.Generate)
	mux.HandleFunc("/api/models", ollamaHandler.ListModels)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
