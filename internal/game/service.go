package game

import (
	"fmt"
	"time"
)

type GameService struct {
	roundLength int
	roundStart  bool
	roundScore  RoundScore
}

func NewGameService() *GameService {
	return &GameService{}
}

/**
* Starts timer for game round.
**/
func (s *GameService) StartRoundService() {

	s.roundStart = true

	go func() {
		// count to 10 and then stop
		time.Sleep(time.Second * gameRoundLength)

		s.roundStart = false // stop round after 10 seconds
	}()

}

/**
* Stops the round incase goroutine failed and clear score.
**/
func (h *GameService) StopRoundService() RoundScore {
	// stop round
	h.roundStart = false

	result := h.GetResultService()

	// clear score
	h.roundScore = RoundScore{}

	return result
}

/**
* Gets the final result of the round.
**/
func (h *GameService) GetResultService() RoundScore {

	// gather score
	return h.roundScore
}

/**
* Submits an answer for the current round.
**/
func (h *GameService) SubmitAnswerService(answer string) error {

	if answer != "A" && answer != "B" && answer != "C" && answer != "D" {
		return fmt.Errorf("Answer needs to be A, B, C, or D.")
	}

	switch answer {
	case "A":
		h.roundScore.A += 1
	case "B":
		h.roundScore.B += 1
	case "C":
		h.roundScore.C += 1
	case "D":
		h.roundScore.D += 1
	}

	return nil
}
