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

func (r GenAIRepository) QueryAnythingLLM(prompt string) (*AnythingLLMResponse, error) {
	// 包裝請求負載，準備發送到 AnythingLLM
	payload := anythingLLMRequest{
		Message:   prompt, // 如果 AnythingLLM 的參數名稱是 Message
		Mode:      "chat",
		SessionID: "identifier-to-partition-chats-by-external-id",
	}

	// 將負載轉為 JSON 格式
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 建立 HTTP 請求，並加入 Bearer Token
	req, err := http.NewRequest("POST", r.cfg.AnythingLLMURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.cfg.AnythingLLMToken) // 加入 Bearer Token 驗證

	// 發送 HTTP 請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // 確保回應的 Body 在函式結束時關閉

	// 驗證回應是否成功
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("非正常 HTTP 狀態: %s", resp.Status)
	}

	// 準備解析流式回應
	reader := bufio.NewReader(resp.Body)
	var result AnythingLLMResponse
	var fullContent string

	// 逐行讀取流式回應
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			// 如果讀取某行失敗，結束函式並返回錯誤
			return nil, fmt.Errorf("解析流時發生錯誤: %s", err)
		}

		// 解析 JSON，讀取每個流的區塊
		var chunk AnythingLLMResponse
		if err := json.Unmarshal(line, &chunk); err != nil {
			continue // 跳過解析錯誤的區塊
		}

		fullContent += chunk.Content // 累積完整內容

		// 檢查流結束的標記 (假設 AnythingLLM 有類似的字段)
		if chunk.Done {
			result = chunk
			break
		}
	}

	// 將完整內容更新到回應結構中
	result.Content = fullContent

	return &result, nil
}
