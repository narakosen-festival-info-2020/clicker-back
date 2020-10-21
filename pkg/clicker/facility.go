package clicker

import (
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
)

func (data *Data) addFacility(name string, numHold, amount float64) {
	tmp := facility.Generate(name, numHold, amount)
	data.facilities[name] = &tmp
	tmp.CountUp(data)
}

// InitFacility is init facilities and exec count up
func (data *Data) InitFacility() {
	data.addFacility("temp", 1, 100)
}
