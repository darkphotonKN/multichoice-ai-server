package genai

import "time"

type ollamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OllamaResponse struct {
	Model      string    `json:"model"`
	CreatedAt  time.Time `json:"created_at"`
	Message    Message   `json:"message"`
	DoneReason string    `json:"done_reason"`
	Done       bool      `json:"done"`
}

type anythingLLMRequest struct {
	Message   string `json:"message"`
	Mode      string `json:"mode"`
	SessionID string `json:"sessionId"`
}

type AnythingLLMResponse struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GenAiQueryRequest struct {
	Prompt string `json:"prompt"`
}

type GenAIMultiQueryRequest struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
