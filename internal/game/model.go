package game

type RoundScore struct {
	A int
	B int
	C int
	D int
}

type RoundScoreResponse struct {
	Score RoundScore `json:"score"`
}

type SubmitAnswerRequest struct {
	Answer string `json:"answer"`
}
