# Ollama API

A Go-based API wrapper for Ollama, providing a simple and efficient way to interact with Ollama's language models through HTTP endpoints.

## Features

- RESTful API endpoints for text generation
- Support for both streaming and non-streaming responses
- Model management capabilities
- Easy configuration through environment variables
- Comprehensive logging for debugging

## Prerequisites

- Go 1.21 or higher
- Ollama running locally or accessible via network

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/taufiksoleh/ollama-api.git
   cd ollama-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the environment:
   - Copy the example environment file:
     ```bash
     cp .env.example .env
     ```
   - Modify the `.env` file with your settings:
     ```env
     SERVER_ADDRESS=:8080
     OLLAMA_ENDPOINT=http://localhost:11434
     ```

4. Run the server:
   ```bash
   go run main.go
   ```

## API Endpoints

### Generate Text

```http
POST /generate
Content-Type: application/json

{
  "model": "llama2",
  "prompt": "What is artificial intelligence?",
  "stream": false,
  "system": "You are a helpful assistant",
  "template": "optional template"
}
```

#### Streaming Response

To receive streaming responses, set `"stream": true` in your request. The server will send Server-Sent Events (SSE) with partial responses as they become available.

### List Models

```http
GET /models
```

Returns a list of available models with their details.

## Configuration

The application can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|----------|
| SERVER_ADDRESS | The address and port for the API server | :8080 |
| OLLAMA_ENDPOINT | The URL of your Ollama instance | http://localhost:11434 |

## Project Structure

```
.
├── config/      # Configuration management
├── handler/     # HTTP request handlers
├── models/      # Data models and types
├── service/     # Business logic and Ollama integration
├── .env         # Environment configuration
└── main.go      # Application entry point
```

## Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments

- [Ollama](https://ollama.ai) for providing the base LLM functionality