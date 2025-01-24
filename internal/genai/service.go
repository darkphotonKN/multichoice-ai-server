package genai

import "fmt"

type GenAIService struct {
	repo *GenAIRepository
}

func NewGenAIService(repo *GenAIRepository) *GenAIService {
	return &GenAIService{
		repo: repo,
	}
}

// func (s *GenAIService) QueryAIService(prompt string) (*OllamaResponse, error) {
func (s *GenAIService) QueryAIService(prompt string) (*AnythingLLMResponse, error) {
	// return s.repo.QueryOllama(prompt)
	return s.repo.QueryAnythingLLM(prompt)
}

func (s *GenAIService) AnswerMultipleChoice(request GenAIMultiQueryRequest) (*string, error) {

	prompt := fmt.Sprintf("You are a quiz assistant. Your task is to choose the correct answer to the following question. Select only one answer (A, B, C, or D) and respond with the corresponding letter—nothing else. Question: %s choose from one of the answers: %v Please respond with a single letter (A, B, C, or D):", request.Question, request.Answers)

	// llamaRes, err := s.repo.QueryOllama(prompt)
	llamaRes, err := s.repo.QueryAnythingLLM(prompt)

	if err != nil {
		return nil, err
	}

	// validate that the answer given is the number 1, 2, 3 or 4
	// if llamaRes.Message.Content != "1" ||
	// 	llamaRes.Message.Content != "2" ||
	// 	llamaRes.Message.Content != "3" ||
	// 	llamaRes.Message.Content != "4" {
	// 	result := "1"
	// 	return &result, nil
	// }

	return &llamaRes.Content, nil
	// return &llamaRes.Message.Content, nil
}
