package game

import (
	"encoding/json"
	"net/http"
)

const gameRoundLength = 12

type GameHandler struct {
	service     *GameService
	roundLength int
	roundStart  bool
	roundScore  RoundScore
}

func NewGameHandler(service *GameService) *GameHandler {
	return &GameHandler{
		service:     service,
		roundLength: gameRoundLength,
	}
}

func (h *GameHandler) StartRoundHandler(w http.ResponseWriter, r *http.Request) {
	// start timer
	h.service.StartRoundService()

	w.Write([]byte("Game Started"))
}

func (h *GameHandler) EndRoundHandler(w http.ResponseWriter, r *http.Request) {

	// stop round, clear data, and get score
	result := h.service.StopRoundService()

	resultResponse, err := json.Marshal(RoundScoreResponse{
		Score: result,
	})

	if err != nil {

		http.Error(w, "Error when attempting to marshal to json.", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultResponse)
}

func (h *GameHandler) SubmitAnswerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request SubmitAnswerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err := h.service.SubmitAnswerService(request.Answer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Answer submitted"))
}
