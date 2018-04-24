package server

import (
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
	Session string      `json:"sessionID"`
	Player  string      `json:"player"`
	Card    models.Card `json:"card"`
	Stack   int         `json:"stack"`
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
		} else if err := PlayCard(game.CurrentBattle, player, &rq.Card, rq.Stack); err != nil {
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
func PlayCard(b *models.Battle, playerState *models.PlayerState, cardID *models.Card, stack int) error {
	if (b.Turn == true && playerState == b.PlayerOne) || (b.Turn == false && playerState == b.PlayerTwo) {
		return errors.New("Invalid Move: It is not this player's turn at this time")
	}
	player := playerState.Player
	if _, err := cardID.Ref(); err != nil {
		return err
	} else if cardID.Type == models.RoadCard {
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
	} else if cardID.Type == models.TuneUpCard {
		tuneup, _ := cardID.AsTuneUp()
		if tuneup.Target == models.TuneupSelf {
			playerState.Tuneups = append(playerState.Tuneups, models.PlayedCard{
				PlayedBy: player.ID,
				Card:     cardID,
			})
		} else if tuneup.Target == models.TuneupRoad {
			b.CardsInPlay[stack].Tuneups = append(b.CardsInPlay[stack].Tuneups, models.PlayedCard{
				PlayedBy: player.ID,
				Card:     cardID,
			})
		}
		if tuneup.OnActivation != nil {
			tuneup.OnActivation()
		}
	} else if cardID.Type == models.DisasterCard {
		disaster, _ := cardID.AsDisaster()
		if disaster.Target == models.DisasterAll {
			b.Disaster = models.PlayedCard{
				PlayedBy: player.ID,
				Card:     cardID,
			}
		} else if disaster.Target == models.DisasterRoad {
			b.CardsInPlay[stack].Disasters = append(b.CardsInPlay[stack].Disasters, models.PlayedCard{
				PlayedBy: player.ID,
				Card:     cardID,
			})
		}
		if disaster.OnActivation != nil {
			disaster.OnActivation()
		}
	}
	playerState.Hand.DrawCardByID(cardID.ID)
	RecalculateScores(b)
	return nil
}
