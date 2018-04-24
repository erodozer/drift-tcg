package server

import (
	"drift/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

// CardHandler will map all cards available into
func CardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	out := struct {
		Cars      map[string]models.Car      `json:"cars"`
		Disasters map[string]models.Disaster `json:"disasters"`
		Roads     map[string]models.Road     `json:"roads"`
		TuneUps   map[string]models.TuneUp   `json:"tuneups"`
	}{
		Cars:      models.Cars,
		Disasters: models.Disasters,
		Roads:     models.Roads,
		TuneUps:   models.TuneUps,
	}
	log.Print(out)
	json.NewEncoder(w).Encode(out)
}
