package cards

import (
	"drift/internal/models"
)

var Disasters = map[string]models.Disaster{
	"68f53073-25d0-4933-be8b-9ab21194b517": models.Disaster{
		Name:   "Rain",
		Impact: models.DisasterAll,
		Action: func() {

		},
	},
}
