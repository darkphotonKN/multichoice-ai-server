package game

import (
	"fmt"
	"time"
)

type GameService struct {
	roundLength int
	roundStart  bool
	roundScore  RoundScore
	players     map[string]RoundScore
}

func NewGameService() *GameService {
	return &GameService{
		players: make(map[string]RoundScore),
	}
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
func (s *GameService) StopRoundService() RoundScore {
	// stop round
	s.roundStart = false

	result := s.GetResultService()

	// clear score
	s.roundScore = RoundScore{}

	return result
}

/**
* Gets the final result of the round.
**/
func (s *GameService) GetResultService() RoundScore {

	// gather score
	return s.roundScore
}

/**
* Submits an answer for the current round.
**/
func (s *GameService) SubmitAnswerService(answer string, player string) error {
	if !s.roundStart {
		return fmt.Errorf("Round has not started yet.")
	}

	if answer != "A" && answer != "B" && answer != "C" && answer != "D" {
		return fmt.Errorf("Answer needs to be A, B, C, or D.")
	}

	// store it globally

	switch answer {
	case "A":
		s.roundScore.A += 1
	case "B":
		s.roundScore.B += 1
	case "C":
		s.roundScore.C += 1
	case "D":
		s.roundScore.D += 1
	}

	return nil
}
