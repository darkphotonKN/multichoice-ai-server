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
	Player string `json:"player"`
	Answer string `json:"answer"`
}
