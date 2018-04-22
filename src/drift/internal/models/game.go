package models

const (
	Undecided = -1
	Downhill  = 0
	Uphill    = 1
)

type Direction int

type PlayedDisaster struct {
	PlayedBy Player
	Card     Disaster
}

type PlayedTuneup struct {
	PlayedBy Player
	Card     TuneUp
}

// RoadStack represent cards in play
type RoadStack struct {
	Road      Road
	Disasters []PlayedDisaster
	Tuneups   []PlayedTuneup
}

type PlayerState struct {
	Player  Player
	Hand    *Hand
	Tuneups []PlayedTuneup
	Score   int
}

type Battle struct {
	PlayerOne       PlayerState
	PlayerTwo       PlayerState
	BattleDisasters []PlayedDisaster
	CardsInPlay     []RoadStack
	Turn            bool
}

type Game struct {
	PlayerOne Player
	PlayerTwo Player

	CurrentBattle *Battle
	Round         int
	Wins          []bool
	Direction     Direction
}

func createHand(d Deck) Hand {
	h := Hand{}
	return h
}

// Begin will start a new battle
func (g Game) Begin() {
	var handOne *Hand
	var handTwo *Hand
	if g.CurrentBattle != nil {
		// carry over hand from previous battle
		handOne = g.CurrentBattle.PlayerOne.Hand
		handTwo = g.CurrentBattle.PlayerTwo.Hand
	} else {
		// make new hands from the player's decks

		// flip a coin to see who goes first
		// flip a coin to see if the battle starts as downhill or uphill
	}
	g.CurrentBattle = &Battle{
		PlayerOne: PlayerState{
			Player:  g.PlayerOne,
			Hand:    handOne,
			Score:   0,
			Tuneups: []PlayedTuneup{},
		},
		PlayerTwo: PlayerState{
			Player:  g.PlayerOne,
			Hand:    handTwo,
			Score:   0,
			Tuneups: []PlayedTuneup{},
		},
		BattleDisasters: []PlayedDisaster{},
		CardsInPlay:     []RoadStack{},
	}
}

// Finish will end a battle, calculating the score and recording the victor
func (g Game) Finish() {

}
