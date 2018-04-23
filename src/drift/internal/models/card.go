package models

const (
	Corner       = 0
	StraightAway = 1
)

type Card interface{}

type RoadType int

type Road struct {
	ID    string
	Name  string
	Type  RoadType `json:"type"`
	Value int      `json:"value"`
}

type Stats struct {
	Cornering int `json:"cornering"`
	Straight  int `json:"straight"`
}

// Car is a type of card that defines what the player will be driving for the duration of a race
//   They are the purest representation of a player
type Car struct {
	ID       string
	Name     string
	Uphill   Stats `json:"uphill"`
	Downhill Stats `json:"downhill"`
}

const (
	TuneupSelf = TuneUpTarget(0)
	TuneupRoad = TuneUpTarget(1)
)

type TuneUpTarget int
type TuneUpAction func(player *Player, battle *Battle)

type TuneUp struct {
	ID     string
	Name   string
	Target TuneUpTarget `json:"target"`
	Action TuneUpAction `json:"-"`
}

const (
	DisasterRoad = 0
	DisasterAll  = 1
)

type DisasterTarget int
type DisasterAction func()

type Disaster struct {
	ID     string
	Name   string
	Impact DisasterTarget `json:"impact"`
	Action DisasterAction `json:"-"`
}
