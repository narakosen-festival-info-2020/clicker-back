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
	data.addFacility("student", 1, 100)
	data.addFacility("3d-printer", 1.6, 240)
	data.addFacility("senbei-refining", 4, 400)
	data.addFacility("lightning-rod", 17, 3000)
	data.addFacility("router", 40, 10000)
	data.addFacility("compiler", 120, 35000)
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
