package main

import (
	drift "drift/internal"
	"drift/internal/cards"
	"drift/internal/models"
	"drift/internal/server"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/jcuga/golongpoll"
)

var manager *golongpoll.LongpollManager

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		DB:   0, // use default DB
	})
}

// Publishes a reduced version of the game to the clients and redis
func publishState(game *models.Game) []byte {
	dto, _ := json.Marshal(game)
	return dto
}

// CardHandler will map all cards available into
func cardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	out := struct {
		Cars      map[string]*models.Car      `json:"cars"`
		Disasters map[string]*models.Disaster `json:"disasters"`
		Roads     map[string]*models.Road     `json:"roads"`
	}{
		Cars:      cards.Cars,
		Disasters: cards.Disasters,
		Roads:     cards.Roads,
	}
	json.NewEncoder(w).Encode(out)
}

type DeckDTO struct {
	Car       string
	TuneUps   []string
	Disasters []string
	Roads     []string
}

func (d *DeckDTO) toDeck() models.Deck {
	car := cards.Cars[d.Car]
	tuneups := []models.TuneUp{}
	disasters := []models.Disaster{}
	roads := []models.Road{}
	for _, card := range cards.TuneUps {
		tuneups = append(tuneups, card)
	}
	for _, card := range cards.Disasters {
		disasters = append(disasters, card)
	}
	for _, card := range cards.Roads {
		roads = append(roads, card)
	}
	return models.Deck{
		Car:       car,
		Tuneups:   tuneups,
		Disasters: disasters,
		Roads:     roads,
	}
}

type JoinGameRequest struct {
	session string  `json:"sessionID"`
	uuid    string  `json:"uuid"`
	player  string  `json:"playerID"`
	deck    DeckDTO `json:"deck"`
}

// join a game
func joinGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	rq := JoinGameRequest{}
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// validate deck
	deck := rq.deck.toDeck()
	if !deck.IsValid() {
		http.Error(w, "Deck does not meet acceptable criteria for competitive gameplay.", http.StatusBadRequest)
		return
	}
	client := getClient()

	if game, err := drift.GetGameFromSession(rq.session, client); err == redis.Nil {
		// create a new session if one isn't found
		log.Print("session not found, creating a new one")
		game = &models.Game{
			PlayerOne: models.Player{
				Id:   rq.player,
				Deck: deck,
			},
			PlayerTwo: models.Player{},
			Round:     0,
			Wins:      []bool{false, false, false},
		}
		// save session
		if _, err = client.Set(rq.session, game, 0).Result(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		log.Panic("something went wrong ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		// if the player is joining a game they're already in, send them the current state of the game
		// this allows a player to rejoin a game they accidentally disconnected from
		if _, err := game.GetPlayer(rq.player); err == nil {
			out := publishState(game)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
			return
		}
		if game.HasStarted {
			http.Error(w, "Match has already started", http.StatusForbidden)
			return
		}

		game.PlayerTwo = models.Player{
			Id:   rq.player,
			Deck: deck,
		}
		// ready to play
		game.Begin()
		manager.Publish(rq.session, publishState(game))

		// save session
		if _, err = client.Set(rq.session, game, 0).Result(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

type PlayCardRequest struct {
	session string              `json:"sessionID"`
	player  string              `json:"player"`
	card    string              `json:"card"`
	target  server.TargetChoice `json:"target"`
	stack   int                 `json:"stack"`
}

func playCard(w http.ResponseWriter, r *http.Request) {
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
	if game, err := drift.GetGameFromSession(rq.session, client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if !game.HasStarted {
			http.Error(w, "The match has not yet started", http.StatusBadRequest)
			return
		}
		// ignore illegal moves by players not in this match
		player, notFound := game.GetPlayer(rq.player)
		if notFound != nil {
			http.Error(w, notFound.Error(), http.StatusBadRequest)
			return
		}
		// attempt to play the card in the server
		if err := server.PlayCard(game.CurrentBattle, player, rq.card, rq.target, rq.stack); err != nil {
			http.Error(w, notFound.Error(), http.StatusBadRequest)
			return
		}
		// send message with game's new state to the players
		// this is naive for now and can get big, ideally we'll want to just send diffs
		dto := publishState(game)
		manager.Publish(rq.session, dto)
		// write back the state in this as well for validation purposes
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto)
	}
}

// A player decides to fold on their turn
func fold(w http.ResponseWriter, r *http.Request) {

}

func main() {
	rand.Seed(time.Now().UnixNano())

	// This launches a goroutine and creates channels for all the plumbing
	m, err := golongpoll.StartLongpoll(golongpoll.Options{}) // default options
	if err != nil {
		log.Fatal("could not start long-polling coroutine", err)
	}
	manager = m

	// Expose events to browsers
	// See subsection on how to interact with the subscription handler
	http.HandleFunc("/events", manager.SubscriptionHandler)
	http.HandleFunc("/cards", cardHandler)
	http.HandleFunc("/join", joinGame)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/", http.FileServer(http.Dir("web"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
