package clicker

import (
	"fmt"

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

// PurchaseFacility is specify facility purchase
func (data *Data) PurchaseFacility(name string) error {
	ope, check := data.facilities[name]
	if !check {
		return fmt.Errorf("Cannot find %s", name)
	}
	check = ope.Purchase(data)
	if !check {
		return fmt.Errorf("We don't have enough cookies: %f", data.GetCount())
	}
	return nil
}
