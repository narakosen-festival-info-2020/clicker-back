package facility

// JSONData is JSON of Data
type JSONData struct {
	id      int
	Name    string  `json:"name"`
	NumHold int     `json:"num_hold"`
	NumGen  float64 `json:"num_gen"`
	Amount  float64 `json:"amount"`
}

// GetJSON is convert Data to JSONData
func (data *Data) GetJSON() JSONData {
	data.RLock()
	defer data.RUnlock()
	return JSONData{
		id:      data.id,
		Name:    data.name,
		NumHold: data.numHold,
		NumGen:  data.numGen,
		Amount:  data.amount,
	}
}

// ID is return id
func (data *JSONData) ID() int {
	return data.id
}
