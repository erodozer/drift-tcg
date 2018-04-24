package server

import (
	"drift/internal/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

type JoinGameRequest struct {
	Session string         `json:"session"`
	Player  string         `json:"player"`
	Deck    []*models.Card `json:"deck"`
}

// join a game
func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	rq := JoinGameRequest{}
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate deck
	if err := models.IsValidDeck(rq.Deck); err != nil {
		log.Print(err.Error())
		http.Error(w, "Deck does not meet acceptable criteria for competitive gameplay.", http.StatusBadRequest)
		return
	}
	deck, _ := models.ToDeck(rq.Deck)
	client := getClient()

	if game, err := models.GetGameFromSession(rq.Session, client); err == redis.Nil {
		// create a new session if one isn't found
		log.Print("session not found, creating a new one")
		game = &models.Game{
			PlayerOne: &models.Player{
				ID:   rq.Player,
				Deck: deck,
			},
			PlayerTwo: &models.Player{},
			Round:     0,
			Wins:      []models.Result{models.UndecidedWin, models.UndecidedWin, models.UndecidedWin},
		}
		// save session
		out, _ := json.Marshal(game)
		if _, err = client.Set(rq.Session, out, 0).Result(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
		}
	} else if err != nil {
		log.Panic("something went wrong ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if game.GetPlayerByID(rq.Player) != nil {
		// if the player is joining a game they're already in, send them the current state of the game
		// this allows a player to rejoin a game they accidentally disconnected from
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
	} else if game.HasStarted {
		http.Error(w, "Match has already started", http.StatusForbidden)
		return
	} else {
		game.PlayerTwo = &models.Player{
			ID:   rq.Player,
			Deck: deck,
		}
		// ready to play
		game.Begin()
		Manager.Publish(rq.Session, game)

		// save session
		out, _ := json.Marshal(game)
		if _, err = client.Set(rq.Session, out, 0).Result(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}
}
