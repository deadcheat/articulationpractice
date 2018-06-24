package types

// Match manage score for match
type Match struct {
	MatchID   string      `json:"matchID"`
	Total     int         `json:"total"`
	Score     int         `json:"score"`
	Questions []*Question `json:"questions,omitempty"`
	Current   *Question   `json:"current,omitempty"`
	StartedAt int64       `json:"startedAt"`
}

// Summary records results for each match
type Summary struct {
	UserID       string
	TotalSuccess int
	History      []*Match
}

// Question hold question text
type Question struct {
	Text    string
	Success bool
}
