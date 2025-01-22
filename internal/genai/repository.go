package genai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darkphotonKN/go-ollama-chat/config"
)

type GenAIRepository struct {
	cfg *config.Config
}

func NewGenAIRepository(cfg *config.Config) *GenAIRepository {
	return &GenAIRepository{
		cfg: cfg,
	}
}

func (r GenAIRepository) QueryOllama(prompt string) (*OllamaResponse, error) {
	// packaging request payload to send to ollama
	payload := ollamaRequest{
		Model: r.cfg.ModelName,
		Messages: []Message{
			Message{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	// converting payload to json binary
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	// post request to ollama chat api
	resp, err := http.Post(r.cfg.OllamaURL, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // close response body when function ends

	// check for error response and return it to request caller if there is one
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: " + resp.Status)
	}

	// setup type for ollama's response, and io Reader for streamed response
	reader := bufio.NewReader(resp.Body)
	var result OllamaResponse
	var fullContent string

	// read stream from ollama
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			// error reading certain line, end function and send error response
			return nil, fmt.Errorf("Error when parsing stream: %s", err)
		}

		// unmarshal ollama respsonse from line into the chunk
		var chunk OllamaResponse

		if err := json.Unmarshal(line, &chunk); err != nil {
			continue // skip current loop
		}

		fullContent += chunk.Message.Content

		// ollama provides "Done" field to indicate end of stream
		if chunk.Done {
			result = chunk
			break
		}
	}

	// replace content with entire content
	result.Message.Content = fullContent

	return &result, nil
}
