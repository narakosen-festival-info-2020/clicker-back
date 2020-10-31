package status

// JSONCore is Core of JSON
type JSONCore struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// GetJSON is convert Core to JSONCore
func (data *Core) GetJSON() JSONCore {
	data.RLock()
	defer data.RUnlock()
	return JSONCore{
		Name:  data.name,
		Value: data.value,
	}
}

// JSONData is Data of JSON
type JSONData struct {
	General []JSONCore `json:"general"`
	Click   []JSONCore `json:"click"`
}

// GetJSON is convert Data to JSONCore
func (data *Data) GetJSON() JSONData {
	ret := JSONData{
		General: make([]JSONCore, len(data.general)),
		Click:   make([]JSONCore, len(data.click)),
	}
	for idx, tmp := range data.general {
		ret.General[idx] = tmp.GetJSON()
	}
	for idx, tmp := range data.click {
		ret.Click[idx] = tmp.GetJSON()
	}
	return ret
}
