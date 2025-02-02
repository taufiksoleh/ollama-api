package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/taufiksoleh/ollama-api/models"
)

type OllamaService struct {
	endpoint string
}

func NewOllamaService(endpoint string) *OllamaService {
	return &OllamaService{endpoint: endpoint}
}

func (s *OllamaService) Generate(req *models.GenerateRequest) (*models.GenerateResponse, chan *models.GenerateResponse, error) {
	url := fmt.Sprintf("%s/api/generate", s.endpoint)

	// Log the request
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Log the complete request body
	log.Printf("Sending generate request to %s with body: %s", url, string(reqBody))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error sending generate request: %v", err)
		return nil, nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		log.Printf("Received error response with status %d: %s", resp.StatusCode, string(body))
		return nil, nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	if req.Stream {
		stream := make(chan *models.GenerateResponse, 100)
		go func() {
			defer resp.Body.Close()
			defer close(stream)

			decoder := json.NewDecoder(resp.Body)
			for {
				var response models.GenerateResponse
				if err := decoder.Decode(&response); err != nil {
					if err != io.EOF {
						log.Printf("Error decoding stream response: %v", err)
					}
					break
				}
				stream <- &response
				if response.Done {
					break
				}
			}
		}()
		return nil, stream, nil
	}

	// Non-streaming response handling
	defer resp.Body.Close()
	var response models.GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error decoding generate response: %v", err)
		return nil, nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Log the successful response
	responseBody, _ := json.Marshal(response)
	log.Printf("Successfully generated response: %s", string(responseBody))

	return &response, nil, nil
}

func (s *OllamaService) ListModels() (*models.ListModelsResponse, error) {
	url := fmt.Sprintf("%s/api/tags", s.endpoint)

	// Log the request
	log.Printf("Fetching models list from %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching models list: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Received error response with status %d: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var response models.ListModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error decoding models list response: %v", err)
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Log the successful response
	log.Printf("Successfully retrieved %d models", len(response.Models))

	return &response, nil
}
