package models

const (
	DisasterRoad = DisasterTarget(0)
	DisasterAll  = DisasterTarget(1)
)

type DisasterTarget int

type Disaster struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Impact DisasterTarget `json:"impact"`
	Action func(points int, player *Player, battle *PlayedRoad) int
}

func fallback(points int, player *Player, battle *PlayedRoad) int {
	return points
}

var Disasters = map[string]*Disaster{
	"68f53073-25d0-4933-be8b-9ab21194b517": &Disaster{
		ID:     "68f53073-25d0-4933-be8b-9ab21194b517",
		Name:   "Rain",
		Impact: DisasterAll,
		Action: fallback,
	},
}
