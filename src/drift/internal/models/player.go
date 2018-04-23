package models

type Player struct {
	ID   string `json:"id"`
	Deck Deck   `json:"deck"`
	Hand Hand   `json:"hand"`
}
