package game

import (
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
* Gets the final result of the round and resets it after.
**/
func (h *GameService) GetResultService() RoundScore {

	// gather score
	return h.roundScore
}
