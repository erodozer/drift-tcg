package server

import (
	"github.com/go-redis/redis"
	"github.com/jcuga/golongpoll"
)

// Game state handling on the server with Redis
// includes
//	 turn order management
//   win conditions
//   user scores
//   user decks
//   user hands
//

var Manager *golongpoll.LongpollManager
var RedisOptions *redis.Options

func getClient() *redis.Client {
	return redis.NewClient(RedisOptions)
}
