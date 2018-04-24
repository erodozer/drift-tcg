package models

type Stats struct {
	Cornering int `json:"cornering"`
	Straight  int `json:"straight"`
}

// Car is a type of card that defines what the player will be driving for the duration of a race
//   They are the purest representation of a player
type Car struct {
	ID       string
	Name     string `json:"name"`
	Uphill   Stats  `json:"uphill"`
	Downhill Stats  `json:"downhill"`
}

// Cars - all available cars that can be selected from when building a deck
var Cars = map[string]Car{
	"b8527474-f0fe-45e5-9fcc-d086a2e64ac5": Car{
		ID:   "b8527474-f0fe-45e5-9fcc-d086a2e64ac5",
		Name: "AE86",
		Uphill: Stats{
			Cornering: 3,
			Straight:  1,
		},
		Downhill: Stats{
			Cornering: 2,
			Straight:  2,
		},
	},
	"4ff6a823-983f-4800-a826-bb688f1c3847": Car{
		ID:   "4ff6a823-983f-4800-a826-bb688f1c3847",
		Name: "R32",
		Uphill: Stats{
			Cornering: 1,
			Straight:  3,
		},
		Downhill: Stats{
			Cornering: 0,
			Straight:  4,
		},
	},
	"5ded0242-f652-4513-9979-e050cc29ce2c": Car{
		ID:   "5ded0242-f652-4513-9979-e050cc29ce2c",
		Name: "FC3S",
		Uphill: Stats{
			Cornering: 2,
			Straight:  2,
		},
		Downhill: Stats{
			Cornering: 1,
			Straight:  3,
		},
	},
	"be2911be-5efd-494a-8bfb-4771ec694a6d": Car{
		ID:   "be2911be-5efd-494a-8bfb-4771ec694a6d",
		Name: "FD3S",
		Uphill: Stats{
			Cornering: 2,
			Straight:  2,
		},
		Downhill: Stats{
			Cornering: 1,
			Straight:  3,
		},
	},
	"175592fc-b7cc-41b5-adcb-fdde3983e821": Car{
		ID:   "175592fc-b7cc-41b5-adcb-fdde3983e821",
		Name: "EVO",
		Uphill: Stats{
			Cornering: 2,
			Straight:  2,
		},
		Downhill: Stats{
			Cornering: 2,
			Straight:  2,
		},
	},
	"f02bd6a1-542e-4ae1-81d2-6dc31d4175ea": Car{
		ID:   "f02bd6a1-542e-4ae1-81d2-6dc31d4175ea",
		Name: "EG6",
		Uphill: Stats{
			Cornering: 3,
			Straight:  1,
		},
		Downhill: Stats{
			Cornering: 1,
			Straight:  3,
		},
	},
}
