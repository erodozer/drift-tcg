package server

import (
	"drift/internal/cards"
	"drift/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	RoadTarget   = TargetChoice(0) // target hi
	PlayerTarget = TargetChoice(1)
	EnvTarget    = TargetChoice(2)
	AddToStack   = TargetChoice(3)
)

type TargetChoice int

type PlayCardRequest struct {
	Session string `json:"sessionID"`
	Player  string `json:"player"`
	Card    string `json:"card"`
}

func PlayCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	rq := PlayCardRequest{}
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client := getClient()
	if game, err := models.GetGameFromSession(rq.Session, client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if !game.HasStarted {
			http.Error(w, "The match has not yet started", http.StatusBadRequest)
			return
		}
		// ignore illegal moves by players not in this match
		if player := game.CurrentBattle.GetPlayerStateByID(rq.Player); player == nil {
			http.Error(w, "Player is not in this match", http.StatusBadRequest)
			return
		} else if err := PlayCard(game.CurrentBattle, player, rq.Card); err != nil {
			// attempt to play the card in the server
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// send message with game's new state to the players
		// this is naive for now and can get big, ideally we'll want to just send diffs
		Manager.Publish(rq.Session, game)
		// write back the state in this as well for validation purposes
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
	}
}

// Plays a card into the game
func PlayCard(b *models.Battle, playerState *models.PlayerState, cardID string) error {
	if (b.Turn == true && playerState == b.PlayerOne) || (b.Turn == false && playerState == b.PlayerTwo) {
		return errors.New("Invalid Move: It is not this player's turn at this time")
	}
	player := playerState.Player
	if _, notFound := models.Roads[cardID]; !notFound {
		b.CardsInPlay = append(
			b.CardsInPlay,
			models.PlayedRoad{
				PlayedCard: models.PlayedCard{
					PlayedBy: player.ID,
					Card:     cardID,
				},
				Tuneups:   []models.PlayedCard{},
				Disasters: []models.PlayedCard{},
			},
		)
		player.Hand.UseCard(cardID)
	}
	if card, notFound := models.TuneUps[cardID]; !notFound {
		tuneup := models.PlayedCard{
			Card:     cardID,
			PlayedBy: player.ID,
		}
		if card.Target == models.TuneupSelf {
			playerState.Tuneups = append(playerState.Tuneups, tuneup)
			player.Hand.UseCard(cardID)
		} else if card.Target == models.TuneupRoad {
			if len(b.CardsInPlay) == 0 {
				return errors.New("Invalid Move: No roads have been played")
			}
			// apply tuneup to most recently played road
			road := b.CardsInPlay[len(b.CardsInPlay)-1]
			road.Tuneups = append(road.Tuneups, tuneup)
			player.Hand.UseCard(cardID)
		}
	}
	if card, notFound := models.Disasters[cardID]; !notFound {
		// apply tuneup to
		disaster := models.PlayedCard{
			Card:     cardID,
			PlayedBy: player.ID,
		}
		if card.Impact == models.DisasterAll {
			b.Disasters = append(b.Disasters, disaster)
			player.Hand.UseCard(cardID)
		} else if card.Impact == models.DisasterRoad {
			if len(b.CardsInPlay) == 0 {
				return errors.New("Invalid Move: No roads have been played")
			}
			road := b.CardsInPlay[len(b.CardsInPlay)-1]
			road.Disasters = append(road.Disasters, disaster)
			player.Hand.UseCard(cardID)
		}
	}

	// recalculate the entire battle state

	return nil
}
