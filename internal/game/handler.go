package game

import (
	"fmt"
	"time"
)

const gameRoundLength = 10

type GameHandler struct {
	repo        *GameService
	roundLength int
	roundScore  roundScore
}

func NewGameHandler(repo *GameService) *GameHandler {
	return &GameHandler{
		repo:        repo,
		roundLength: gameRoundLength,
	}
}

/**
* Starts timer for game round.
**/
func (h *GameHandler) StartRound() {
	ticker := time.NewTicker(time.Duration(h.roundLength))

	count := <-ticker.C

	fmt.Println(count)
}

/**
* Get final result of the round reset round.
**/
func (h *GameHandler) GetResult() roundScore {
	return h.roundScore
}
