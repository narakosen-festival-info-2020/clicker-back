package clicker

import (
	"fmt"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/status"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
)

// JSONData is JSON of data (count etc)
type JSONData struct {
	Count float64 `json:"count"`
}

// GetJSON is convert data to JSON
func (data *Data) GetJSON() JSONData {
	return JSONData{
		Count: data.GetCount(),
	}
}

// GetFacilityJSON is get spefific facility
func (data *Data) GetFacilityJSON(name string) (facility.JSONData, error) {
	if val, check := data.facilities[name]; check {
		return val.GetJSON(), nil
	}
	return facility.JSONData{}, fmt.Errorf("Cannot find %s", name)
}

// GetAllFacilityJSON is get all facility
func (data *Data) GetAllFacilityJSON() []facility.JSONData {
	ret := []facility.JSONData{}
	for _, val := range data.facilities {
		ret = append(ret, val.GetJSON())
	}
	return ret
}

// JSONStatements is convert statements to JSON
type JSONStatements struct {
	Statements status.JSONData `json:"statements"`
}

// GetStatements is get all statements
func (data *Data) GetStatements() JSONStatements {
	return JSONStatements{
		Statements: data.statements.GetJSON(),
	}
}
