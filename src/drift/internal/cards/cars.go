package cards

import (
	"drift/internal/models"
)

// Cars - all available cars that can be selected from when building a deck
var Cars = map[string]models.Car{
	"b8527474-f0fe-45e5-9fcc-d086a2e64ac5": models.Car{
		Name: "AE86",
		Uphill: models.Stats{
			Cornering: 3,
			Straight:  1,
		},
		Downhill: models.Stats{
			Cornering: 2,
			Straight:  2,
		},
	},
	"4ff6a823-983f-4800-a826-bb688f1c3847": models.Car{
		Name: "R32",
		Uphill: models.Stats{
			Cornering: 1,
			Straight:  3,
		},
		Downhill: models.Stats{
			Cornering: 0,
			Straight:  4,
		},
	},
	"5ded0242-f652-4513-9979-e050cc29ce2c": models.Car{
		Name: "FC3S",
		Uphill: models.Stats{
			Cornering: 2,
			Straight:  2,
		},
		Downhill: models.Stats{
			Cornering: 1,
			Straight:  3,
		},
	},
	"be2911be-5efd-494a-8bfb-4771ec694a6d": models.Car{
		Name: "FD3S",
		Uphill: models.Stats{
			Cornering: 2,
			Straight:  2,
		},
		Downhill: models.Stats{
			Cornering: 1,
			Straight:  3,
		},
	},
	"175592fc-b7cc-41b5-adcb-fdde3983e821": models.Car{
		Name: "EVO",
		Uphill: models.Stats{
			Cornering: 2,
			Straight:  2,
		},
		Downhill: models.Stats{
			Cornering: 2,
			Straight:  2,
		},
	},
	"f02bd6a1-542e-4ae1-81d2-6dc31d4175ea": models.Car{
		Name: "EG6",
		Uphill: models.Stats{
			Cornering: 3,
			Straight:  1,
		},
		Downhill: models.Stats{
			Cornering: 1,
			Straight:  3,
		},
	},
}
