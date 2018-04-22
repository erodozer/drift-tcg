package cards

import (
	"drift/internal/models"
)

var TuneUps = map[string]models.TuneUp{
	"2de2704e-4b42-43e4-9e68-9919ca3ea110": models.TuneUp{
		Name:   "Supercharger",
		Target: models.TuneupSelf,
		Action: func(player *models.Player, battle *models.Battle) {

		},
	},
}
