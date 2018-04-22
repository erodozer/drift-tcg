package models

import "math/rand"

type Deck struct {
	Car       Car
	Tuneups   []TuneUp
	Roads     []Road
	Disasters []Disaster
}

type Hand struct {
	Cards []Card
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

// Generate a hand randomly from the deck
func (d Deck) createHand() Hand {
	h := Hand{}
	cards := append([]Card(nil), d.Car)
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
