package facility

// JSONData is JSON of Data
type JSONData struct {
	Name    string  `json:"name"`
	NumHold float64 `json:"num_hold"`
	NumGen  float64 `json:"num_gen"`
	Amount  float64 `json:"amount"`
}

// GetJSON is convert Data to JSONData
func (data *Data) GetJSON() JSONData {
	data.RLock()
	defer data.RUnlock()
	return JSONData{
		Name:    data.name,
		NumHold: data.numHold,
		NumGen:  data.numGen,
		Amount:  data.amount,
	}
}
