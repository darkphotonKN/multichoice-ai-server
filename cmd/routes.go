package main

import (
	"net/http"

	"github.com/darkphotonKN/go-ollama-chat/config"
	"github.com/darkphotonKN/go-ollama-chat/internal/game"
	"github.com/darkphotonKN/go-ollama-chat/internal/genai"
	"github.com/rs/cors"
)

func SetupRoutes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	// --- general settings --
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// --- Gen AI ---

	// -- setup --
	genAIRepo := genai.NewGenAIRepository(cfg)
	genAIService := genai.NewGenAIService(genAIRepo)
	genAIHandler := genai.NewGenAIHandler(genAIService)

	// -- api routes --
	mux.HandleFunc("/api/query", genAIHandler.QueryAIHandler)
	mux.HandleFunc("/api/query-multi-choice", genAIHandler.QueryAIHandler)

	// --- Game ---

	// -- setup --
	gameService := game.NewGameService()
	gameHandler := game.NewGameHandler(gameService)

	// -- api routes --
	mux.HandleFunc("/api/game/start", gameHandler.StartRoundHandler)
	mux.HandleFunc("/api/game/end", gameHandler.EndRoundHandler)
	mux.HandleFunc("/api/game/submitAnswer", gameHandler.SubmitAnswerHandler)

	handler := c.Handler(mux)

	return handler
}
