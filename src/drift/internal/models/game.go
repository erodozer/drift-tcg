package models

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/go-redis/redis"
)

type Direction bool

const (
	Downhill = Direction(false)
	Uphill   = Direction(true)
)

type PlayedCard struct {
	PlayedBy string `json:"playedBy"`
	Card     *Card  `json:"card"`
}

// RoadStack represent cards in play
type PlayedRoad struct {
	PlayedCard
	Disasters []PlayedCard `json:"disasters"`
	Tuneups   []PlayedCard `json:"tuneups"`
}

// Get the calculated worth of this road tile for the player
func (p *PlayedRoad) Value(player *Player) int {
	card, _ := p.Card.AsRoad()
	points := card.Value
	for _, card := range p.Tuneups {
		tuneup, _ := card.Card.AsTuneUp()
		points = tuneup.OnScoring(points, player)
	}
	for _, card := range p.Disasters {
		disaster, _ := card.Card.AsTuneUp()
		points = disaster.OnScoring(points, player)
	}
	return points
}

type PlayerState struct {
	Player  *Player      `json:"player"`
	Pile    *CardStack   `json:"pile"`
	Discard *CardStack   `json:"discarded"`
	Hand    *CardStack   `json:"hand"`
	Tuneups []PlayedCard `json:"tuneups"`
	Score   int          `json:"score"`
}

// DrawCard removes a card from the player's pile and inserts it into their hand
func (p *PlayerState) DrawCard() {
	p.Hand.AddCard(p.Pile.DrawCard())
}

type Battle struct {
	PlayerOne   *PlayerState
	PlayerTwo   *PlayerState
	Disaster    PlayedCard
	CardsInPlay []PlayedRoad
	Direction   Direction
	StartTurn   bool
	Turn        bool
}

type Result int
type Style int

const (
	UndecidedWin = Result(0)
	PlayerOneWin = Result(1)
	PlayerTwoWin = Result(2)
	TimeAttack   = Style(0)
	SuddenDeath  = Style(1)
)

type Game struct {
	PlayerOne *Player `json:"playerOne"`
	PlayerTwo *Player `json:"playerTwo"`

	Style         Style    `json:"style"`
	CurrentBattle *Battle  `json:"currentBattle"`
	Round         int      `json:"round"`
	Wins          []Result `json:"wins"`
	HasStarted    bool     `json:"hasStarted"`
}

// Begin will start a new battle
func (g *Game) Begin() {
	var turn bool
	var direction Direction
	if g.CurrentBattle != nil {
		// carry over hand from previous battle
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
			Player:  g.PlayerOne,
			Score:   0,
			Tuneups: []PlayedCard{},
			Pile:    g.PlayerOne.Deck,
			Hand:    g.PlayerOne.Deck.CreateHand(g),
			Discard: &CardStack{},
		},
		PlayerTwo: &PlayerState{
			Player:  g.PlayerTwo,
			Score:   0,
			Tuneups: []PlayedCard{},
		},
		Disaster:    PlayedCard{},
		CardsInPlay: []PlayedRoad{},
		Direction:   direction,
		StartTurn:   turn,
		Turn:        turn,
	}
	g.HasStarted = true
}

// Finish will end a battle, calculating the score and recording the victor
func (g Game) Finish() {

}

func (g Game) GetPlayerByID(id string) *Player {
	if g.PlayerOne.ID == id {
		return g.PlayerOne
	} else if g.PlayerTwo.ID == id {
		return g.PlayerTwo
	} else {
		return nil
	}
}

func (g Battle) GetPlayerStateByID(id string) *PlayerState {
	if g.PlayerOne.Player.ID == id {
		return g.PlayerOne
	} else if g.PlayerTwo.Player.ID == id {
		return g.PlayerTwo
	} else {
		return nil
	}
}

func (g Battle) GetPlayerState(player *Player) *PlayerState {
	if g.PlayerOne.Player == player {
		return g.PlayerOne
	} else if g.PlayerTwo.Player == player {
		return g.PlayerTwo
	} else {
		return nil
	}
}

func GetGameFromSession(session string, client *redis.Client) (*Game, error) {
	if data, err := client.Get(session).Result(); err != nil {
		return nil, err
	} else {
		game := Game{}
		if err := json.Unmarshal([]byte(data), &game); err != nil {
			return nil, errors.New("Could not deserialize game from storage")
		}
		return &game, nil
	}
}
