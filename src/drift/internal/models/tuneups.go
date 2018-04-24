package models

const (
	TuneupSelf = TuneUpTarget(0)
	TuneupRoad = TuneUpTarget(1)
)

type TuneUpTarget int

type TuneUp struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Target TuneUpTarget `json:"target"`
	// activates upon playing the card
	OnActivation func() `json:"-"`
	// activates upon calculating the score
	OnScoring func(points int, player *Player) int `json:"-"`
}

var TuneUps = map[string]TuneUp{
	"2de2704e-4b42-43e4-9e68-9919ca3ea110": TuneUp{
		ID:     "2de2704e-4b42-43e4-9e68-9919ca3ea110",
		Name:   "Supercharger",
		Target: TuneupSelf,
		OnScoring: func(points int, player *Player) int {
			return points
		},
	},
}
