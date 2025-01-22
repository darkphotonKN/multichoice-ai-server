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
