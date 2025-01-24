package genai

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GenAIHandler struct {
	service *GenAIService
}

func NewGenAIHandler(service *GenAIService) *GenAIHandler {
	return &GenAIHandler{
		service: service,
	}
}

func (h *GenAIHandler) QueryAIHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request payload: %v", r.Body)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request GenAiQueryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	queryResponse, err := h.service.QueryAIService(request.Prompt)

	if err != nil {
		log.Printf("Error querying AI: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(queryResponse)

	fmt.Printf("Error when attempting to marshal query response: %s\n", err)

	w.Write(response)
}

func (h *GenAIHandler) AnswerMultipleChoice(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request GenAIMultiQueryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	queryResponse, err := h.service.AnswerMultipleChoice(request)

	if err != nil {
		log.Printf("Error querying AI: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(queryResponse)

	fmt.Printf("Error when attempting to marshal query response: %s\n", err)

	w.Write(response)
}
