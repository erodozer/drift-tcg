package cards

func GetCard(id string) interface{} {
	if card, err := Cars[id]; !err {
		return card
	}
	if card, err := TuneUps[id]; !err {
		return card
	}
	if card, err := Disasters[id]; !err {
		return card
	}
	if card, err := Roads[id]; !err {
		return card
	}
	return nil
}
