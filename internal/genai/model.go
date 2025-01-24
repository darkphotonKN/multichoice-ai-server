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
	Error        string   `json:"error"`
	TextResponse string   `json:"textResponse"`
	Sources      []Source `json:"sources"`
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

type Source struct {
	Text        string `json:"text"`
	Id          string `json:"id"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	DocAuthor   string `json:"docAuthor"`
	Description string `json:"description"`
	DocSource   string `json:"docSource"`
	Published   string `json:"published"`
	WordCount   int    `json:"wordCount"`
	TokenCount  int    `json:"tokenCount"`
}
