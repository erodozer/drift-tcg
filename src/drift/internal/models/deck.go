package models

import (
	"errors"
	"math/rand"
)

type CardStack struct {
	Car   *Card   `json:"car"`
	Cards []*Card `json:"cards"`
}

func ToDeck(cards []*Card) (*CardStack, error) {
	for index, card := range cards {
		if _, err := card.AsCar(); err == nil {
			return &CardStack{
				Car:   card,
				Cards: append(cards[:index], cards[index+1:]...),
			}, nil
		}
	}
	return nil, errors.New("Car not found in deck")
}

// IsValid goes over the provided cards and makes sure they don't extend past the limits of play
func IsValidDeck(cards []*Card) error {
	if len(cards) < 11 {
		return errors.New("Not Enough Cards in Deck")
	} else if len(cards) > 56 {
		return errors.New("Too Many Cards")
	}
	cars, tuneups, roads, disasters := 0, 0, 0, 0
	for _, card := range cards {
		if _, err := card.Ref(); err != nil {
			return err
		} else {
			switch card.Type {
			case TuneUpCard:
				tuneups++
			case RoadCard:
				roads++
			case DisasterCard:
				disasters++
			case CarCard:
				cars++
			}
		}
	}
	if tuneups > 10 {
		return errors.New("Too Many Cards: TuneUps")
	}
	if roads < 10 {
		return errors.New("Not Enough Cards: Roads")
	} else if roads > 35 {
		return errors.New("Too Many Cards: Roads")
	}
	if disasters > 10 {
		return errors.New("Too Many Cards: Disasters")
	}
	if cars > 1 {
		return errors.New("Too Many Cards: Car")
	} else if cars < 1 {
		return errors.New("Not Enough Cards: Car")
	}
	return nil
}

func (h *CardStack) AddCard(card *Card) {
	h.Cards = append(h.Cards, card)
}

// pops a card from the stack
func (h *CardStack) DrawCard() *Card {
	index := rand.Intn(len(h.Cards))
	return h.DrawCardAtIndex(index)
}

// pops a card out of the stack by its ID
func (h *CardStack) DrawCardByID(ID string) *Card {
	for index, card := range h.Cards {
		if card.ID == ID {
			return h.DrawCardAtIndex(index)
		}
	}
	return nil
}

// pops a card from the stack at a specific position
func (h *CardStack) DrawCardAtIndex(index int) *Card {
	card := h.Cards[index]
	h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)
	return card
}

// Generate a hand randomly from the deck
func (d *CardStack) CreateHand(game *Game) *CardStack {
	cards := []*Card{}
	added := 0
	var draw int
	switch game.Style {
	case SuddenDeath:
		draw = 10
	case TimeAttack:
		draw = 7
	default:
		draw = 7
	}
	tmpStack := CardStack{
		Car:   &Card{},
		Cards: d.Cards[:len(d.Cards)],
	}
	for added < draw {
		cards = append(cards, tmpStack.DrawCard())
		added++
	}
	return &CardStack{
		Car:   d.Car,
		Cards: cards,
	}
}
