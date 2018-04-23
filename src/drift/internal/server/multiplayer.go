package server

import (
	"drift/internal/cards"
	"drift/internal/models"
	"errors"
)

const (
	RoadTarget   = TargetChoice(0) // target hi
	PlayerTarget = TargetChoice(1)
	EnvTarget    = TargetChoice(2)
	AddToStack   = TargetChoice(3)
)

type TargetChoice int

// Plays a card into the game
func PlayCard(b *models.Battle, player *models.PlayerState, cardID string) error {
	if (b.Turn == true && player == b.PlayerOne) || (b.Turn == false && player == b.PlayerTwo) {
		return errors.New("Invalid Move: It is not this player's turn at this time")
	}
	player.Hand.UseCard(cardID)
	if card, notFound := cards.Roads[cardID]; !notFound {
		b.CardsInPlay = append(
			b.CardsInPlay,
			&models.RoadStack{
				Road:     card,
				PlayedBy: player.Player,
				Tuneups:  []*models.PlayedTuneup{},
			},
		)
		player.Hand.UseCard(cardID)
	}
	if card, notFound := cards.TuneUps[cardID]; !notFound {
		if card.Target == models.TuneupSelf {
			player.Hand.UseCard(cardID)
		} else if card.Target == models.TuneupRoad {

		}
	}
	if card, notFound := cards.Disasters[cardID]; !notFound {

	}

	// recalculate the entire battle state

	return nil
}
