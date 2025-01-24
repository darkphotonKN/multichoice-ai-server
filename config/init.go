package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OllamaURL        string
	ModelName        string
	Port             string
	AnythingLLMURL   string
	AnythingLLMToken string
}

func LoadConfig() Config {
	// load envs
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("asdsad ", err)
	}

	// return config object
	return Config{
		OllamaURL:        getEnv("OLLAMA_URL", "http://localhost:11434/api/chat"),
		ModelName:        getEnv("OLLAMA_MODEL", "llama2"),
		Port:             getEnv("PORT", "8080"),
		AnythingLLMURL:   getEnv("ANYTHING_LLM_URL", "http://localhost:3001/api/v1/workspace/cloud-interactive/chat"),
		AnythingLLMToken: getEnv("Anything_LLM_Token", "H144KFF-N364DMD-G3XRSMR-NFQKASB"),
	}
}

// simple helper for env fallback
func getEnv(key, fallback string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback

}
