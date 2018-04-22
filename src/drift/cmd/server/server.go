package main

import (
	"drift/internal/cards"
	"drift/internal/models"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// CardHandler will map all cards available into
func CardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	out := struct {
		Cars      map[string]models.Car      `json:"cars"`
		Disasters map[string]models.Disaster `json:"disasters"`
		Roads     map[string]models.Road     `json:"roads"`
	}{
		Cars:      cards.Cars,
		Disasters: cards.Disasters,
		Roads:     cards.Roads,
	}
	json.NewEncoder(w).Encode(out)
}

// join a game
func JoinGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	session := r.FormValue("session-id")
	playerId := r.FormValue("uuid")
	if playerId == "" {
		playerId = session
	}
	var deck models.Deck
	err := json.Unmarshal([]byte(r.FormValue("deck")), deck)
	// validate deck
	if !deck.IsValid() {
		http.Error(w, "Deck does not meet acceptable criteria for competitive gameplay.", http.StatusBadRequest)
		return
	}
	client := getClient()
	data, err := client.Get(session).Result()
	// create a new session if one isn't found
	if err == redis.Nil {
		log.Print("session not found, creating a new one")
		game := models.Game{
			PlayerOne: models.Player{
				Id:    playerId,
				Deck:  deck,
				Hand:  models.Hand{},
				Score: 0,
			},
			PlayerTwo: nil,
			Round:     0,
			Wins:      []bool{false, false, false},
			Direction: -1,
		}
		r.Set(session)
	} else if err != nil {
		log.Panic("something went wrong ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		game := models.Game{}
		err := json.Unmarshal([]byte(data), &game)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if game.PlayerTwo != nil {
			http.Error(w, "Match has already started", http.StatusForbidden)
			return
		}
		game.PlayerTwo = models.Player{
			Id:    playerId,
			Deck:  deck,
			Hand:  models.Hand{},
			Score: 0,
		}
	}
	// ready to play
	if game.PlayerOne != nil && game.PlayerTwo != nil {
		// deal the hands

	}

	// save session
	_, err = r.Set(session, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := http.Response{
		status: http.StatusOK,
	}
	response.Write(w)
}

// Long poll the server for events
func Events(w http.ResponseWriter, r *http.Request) {

}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/cards", CardHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
