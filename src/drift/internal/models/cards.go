package models

import "errors"

type CardType string

const (
	CarCard      = CardType("car")
	TuneUpCard   = CardType("tuneup")
	DisasterCard = CardType("disaster")
	RoadCard     = CardType("road")
)

type Card struct {
	ID   string   `json:"id"`
	Type CardType `json:"_type"`
}

func (c *Card) AsRoad() (*Road, error) {
	if card, found := Roads[c.ID]; found {
		return &card, nil
	}
	return nil, errors.New("Road card does not exist: Invalid ID " + c.ID)
}

func (c *Card) AsTuneUp() (*TuneUp, error) {
	if card, found := TuneUps[c.ID]; found {
		return &card, nil
	}
	return nil, errors.New("Tuneup card does not exist: Invalid ID " + c.ID)
}

func (c *Card) AsDisaster() (*Disaster, error) {
	if card, found := Disasters[c.ID]; found {
		return &card, nil
	}
	return nil, errors.New("Disaster card does not exist: Invalid ID " + c.ID)
}

func (c *Card) AsCar() (*Car, error) {
	if card, found := Cars[c.ID]; found {
		return &card, nil
	}
	return nil, errors.New("Car card does not exist: Invalid ID " + c.ID)
}

// Ref gets the actual functional reference to a card by its ID
func (c *Card) Ref() (interface{}, error) {
	switch c.Type {
	case RoadCard:
		return c.AsRoad()
	case TuneUpCard:
		return c.AsTuneUp()
	case DisasterCard:
		return c.AsDisaster()
	case CarCard:
		return c.AsCar()
	default:
		return nil, errors.New("Reference card does not exist: Invalid Type")
	}
	return nil, errors.New("Reference card does not exist: Invalid ID " + c.ID + ", " + string(c.Type))
}
