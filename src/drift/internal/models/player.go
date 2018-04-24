package models

type Player struct {
	ID   string     `json:"id"`
	Deck *CardStack `json:"deck"`
}
