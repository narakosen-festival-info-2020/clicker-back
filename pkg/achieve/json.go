package achieve

// JSONCore is Core of JSON
type JSONCore struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Unlocked bool   `json:"unlocked"`
}

// GetJSON is convert Core to JSONCore
func (data *Core) GetJSON() JSONCore {
	data.RLock()
	defer data.RUnlock()
	return JSONCore{
		ID:       data.id,
		Name:     data.name,
		Unlocked: data.unlocked,
	}
}

// JSONData is Data of JSON
type JSONData struct {
	Achievements []JSONCore `json:"achievements"`
}

// GetJSON is convert Data to JSONCore
func (data *Data) GetJSON() JSONData {
	ret := JSONData{
		Achievements: make([]JSONCore, len(data.achievements)),
	}
	for idx, tmp := range data.achievements {
		ret.Achievements[idx] = tmp.GetJSON()
	}
	return ret
}
