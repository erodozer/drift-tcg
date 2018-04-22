package models

type Deck struct {
	car       Car
	tuneups   []TuneUp
	roads     []Road
	disasters []Disaster
}

type Hand struct {
	car       Car
	tuneups   []TuneUp
	roads     []Road
	disasters []Disaster
}
