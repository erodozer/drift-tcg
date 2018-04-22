package models

import "errors"
import "math/rand"

type Direction bool

const (
	Downhill = Direction(false)
	Uphill   = Direction(true)
)

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
	Road      *Road
	Disasters []*PlayedDisaster
	Tuneups   []*PlayedTuneup
	PlayedBy  *Player
}

type PlayerState struct {
	Player  *Player
	Hand    *Hand
	Tuneups []PlayedTuneup
	Score   int
}

type Battle struct {
	PlayerOne       *PlayerState
	PlayerTwo       *PlayerState
	BattleDisasters []*PlayedDisaster
	CardsInPlay     []*RoadStack
	Direction       Direction
	StartTurn       bool
	Turn            bool
}

type Game struct {
	PlayerOne Player
	PlayerTwo Player

	CurrentBattle *Battle
	Round         int
	Wins          []bool
	HasStarted    bool
}

// Begin will start a new battle
func (g Game) Begin() {
	var handOne *Hand
	var handTwo *Hand
	var turn bool
	var direction Direction
	if g.CurrentBattle != nil {
		// carry over hand from previous battle
		handOne = g.CurrentBattle.PlayerOne.Hand
		handTwo = g.CurrentBattle.PlayerTwo.Hand
		turn = !g.CurrentBattle.StartTurn
		direction = !g.CurrentBattle.Direction
	} else {
		// make new hands from the player's decks

		// flip a coin to see who goes first
		turn = rand.Intn(2) == 1
		// flip a coin to see if the battle starts as downhill or uphill
		direction = rand.Intn(2) == 1
	}
	g.CurrentBattle = &Battle{
		PlayerOne: &PlayerState{
			Player:  &g.PlayerOne,
			Hand:    handOne,
			Score:   0,
			Tuneups: []PlayedTuneup{},
		},
		PlayerTwo: &PlayerState{
			Player:  &g.PlayerTwo,
			Hand:    handTwo,
			Score:   0,
			Tuneups: []PlayedTuneup{},
		},
		BattleDisasters: []*PlayedDisaster{},
		CardsInPlay:     []*RoadStack{},
		Direction:       direction,
		StartTurn:       turn,
		Turn:            turn,
	}
	g.HasStarted = true
}

// Finish will end a battle, calculating the score and recording the victor
func (g Game) Finish() {

}

func (g Game) GetPlayer(id string) (*Player, error) {
	if g.PlayerOne.Id == id {
		return &g.PlayerOne, nil
	} else if g.PlayerTwo.Id == id {
		return &g.PlayerTwo, nil
	} else {
		return nil, errors.New("Player is not playing this match")
	}
}
