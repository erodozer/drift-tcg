package models

import (
	"errors"
	"math/rand"
)

type Deck struct {
	Car       string   `json:"car"`
	Tuneups   []string `json:"tuneups"`
	Roads     []string `json:"roads"`
	Disasters []string `json:"disasters"`
}

type Hand struct {
	Cards []string `json:"cards"` // ids of cards
}

// IsValid goes over the provided cards and makes sure they don't extend past the limits of play
func (d Deck) IsValid() bool {
	tuneups := len(d.Tuneups)
	roads := len(d.Roads)
	disasters := len(d.Disasters)

	return tuneups <= 5 &&
		roads >= 10 &&
		disasters <= 5 &&
		(tuneups+roads+disasters) >= 12
}

func (h Hand) UseCard(id string) error {
	if len(h.Cards) == 0 {
		return errors.New("Hand is empty")
	}
	index := -1
	for i, card := range h.Cards {
		if card == id {
			index = i
			break
		}
	}
	if index > -1 {
		h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)
	} else {
		return errors.New("Card is not in player's hand")
	}
	return nil
}

// Generate a hand randomly from the deck
func (d Deck) createHand() Hand {
	h := Hand{}
	cards := append([]string{}, d.Car)
	added := 0
	for added < 12 {
		choice := rand.Intn(3)
		if choice == 0 {
			if len(d.Tuneups) == 0 {
				continue
			}
			index := rand.Intn(len(d.Tuneups))
			card := d.Tuneups[index]
			d.Tuneups = append(d.Tuneups[:index], d.Tuneups[index+1:]...)
			cards = append(cards, card)
			added++
		} else if choice == 1 {
			if len(d.Roads) == 0 {
				continue
			}
			index := rand.Intn(len(d.Roads))
			card := d.Roads[index]
			d.Roads = append(d.Roads[:index], d.Roads[index+1:]...)
			cards = append(cards, card)
			added++
		} else if choice == 2 {
			if len(d.Disasters) == 0 {
				continue
			}
			index := rand.Intn(len(d.Disasters))
			card := d.Disasters[index]
			d.Disasters = append(d.Disasters[:index], d.Disasters[index+1:]...)
			cards = append(cards, card)
			added++
		}
	}
	h.Cards = cards
	return h
}
