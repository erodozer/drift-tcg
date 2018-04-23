package main

import (
	"drift/internal/server"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/jcuga/golongpoll"
)

// A player decides to fold on their turn
func fold(w http.ResponseWriter, r *http.Request) {

}

func main() {
	rand.Seed(time.Now().UnixNano())

	// This launches a goroutine and creates channels for all the plumbing
	// default options
	if m, err := golongpoll.StartLongpoll(golongpoll.Options{}); err != nil {
		log.Fatal("could not start long-polling coroutine", err)
	} else {
		server.Manager = m
	}

	if options, err := redis.ParseURL(os.Getenv("REDIS_URL")); err != nil {
		log.Fatal("could not parse redis url", err)
	} else {
		server.RedisOptions = options
	}

	// Expose events to browsers
	// See subsection on how to interact with the subscription handler
	http.HandleFunc("/events", server.Manager.SubscriptionHandler)
	http.HandleFunc("/cards", server.CardHandler)
	http.HandleFunc("/join", server.JoinGameHandler)
	http.HandleFunc("/play", server.PlayCardHandler)
	http.Handle("/", http.FileServer(http.Dir("./drift/web/")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
