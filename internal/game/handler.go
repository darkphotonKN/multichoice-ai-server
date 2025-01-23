package game

import (
	"encoding/json"
	"net/http"
)

const gameRoundLength = 10

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

/**
* Get final result of the round reset round.
**/
func (h *GameHandler) GetResultHandler(w http.ResponseWriter, r *http.Request) {
	// gather score
	result := h.service.GetResultService()

	resultResponse, err := json.Marshal(RoundScoreResponse{
		Score: result,
	})

	if err != nil {
		w.Write([]byte("Error when attempting to marshal to json."))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultResponse)
}
