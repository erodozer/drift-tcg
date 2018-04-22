package models

const (
	Undecided = -1
	Downhill  = 0
	Uphill    = 1
)

type Direction int

type PlayedDisaster struct {
	playedBy Player
	card     Disaster
}

type PlayedTuneup struct {
	playedBy Player
	card     TuneUp
}

// RoadStack represent cards in play
type RoadStack struct {
	road      Road
	disasters []PlayedDisaster
	tuneups   []PlayedTuneup
}

type PlayerState struct {
	plyaer  Player
	tuneups []PlayedTuneup
}

type Game struct {
	playerOne       PlayerState
	playerTwo       PlayerState
	CardsInPlay     []RoadStack
	BattleDisasters []PlayedDisaster
	Round           int
	Wins            []bool
	Direction       Direction
}
