package internal

import (
	"drift/internal/models"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis"
)

func GetGameFromSession(session string, client *redis.Client) (*models.Game, error) {
	if data, err := client.Get(session).Result(); err != nil {
		return nil, err
	} else {
		game := models.Game{}
		if err := json.Unmarshal([]byte(data), &game); err != nil {
			return nil, errors.New("Could not deserialize game from storage")
		}
		return &game, nil
	}
}
