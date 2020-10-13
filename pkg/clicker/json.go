package clicker

// JSONData is JSON of data (count etc)
type JSONData struct {
	Count float64 `json:"count"`
}

// GetJSON is convert data to JSON
func (data *Data) GetJSON() JSONData {
	data.RLock()
	defer data.RUnlock()
	return JSONData{
		Count: data.count,
	}
}
