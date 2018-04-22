package models

type Player struct {
	Id    string
	Score int
	Hand  Hand
	Deck  Deck
}
