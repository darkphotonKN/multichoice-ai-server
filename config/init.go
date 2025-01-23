package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OllamaURL string
	ModelName string
	Port      string
}

func LoadConfig() Config {
	// load envs
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("asdsad ", err)
	}

	// return config object
	return Config{
		OllamaURL: getEnv("OLLAMA_URL", "http://localhost:11434/api/chat"),
		ModelName: getEnv("OLLAMA_MODEL", "llama2"),
		Port:      getEnv("PORT", "8080"),
	}
}

// simple helper for env fallback
func getEnv(key, fallback string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback

}
