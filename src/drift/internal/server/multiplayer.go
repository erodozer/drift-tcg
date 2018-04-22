package server

import (
	"drift/internal/cards"
	"drift/internal/models"
	"errors"
)

const (
	RoadTarget   = 0
	PlayerTarget = 1
	AddToStack   = -1
)

type TargetChoice int

// Plays a card into the game
func PlayCard(b *models.Battle, player *models.Player, cardID string, target TargetChoice, stack int) error {
	if (b.Turn == true && player == b.PlayerOne.Player) || (b.Turn == false && player == b.PlayerTwo.Player) {
		return errors.New("Invalid Move: It is not this player's turn at this time")
	}
	if target == RoadTarget {
		if stack > 0 && stack < len(b.CardsInPlay) {
			// play on a stack
		} else if stack == AddToStack {
			// adds a new road tile
			if card, found := cards.Roads[cardID]; found {
				b.CardsInPlay = append(
					b.CardsInPlay,
					&models.RoadStack{
						Road:     &card,
						PlayedBy: player,
						Tuneups:  []*models.PlayedTuneup{},
					},
				)
			} else {
				return errors.New("Invalid Move: Road card with ID does not exists")
			}
		} else {

		}
	}
	return nil
}
