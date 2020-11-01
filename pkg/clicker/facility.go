package clicker

import (
	"fmt"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
)

func (data *Data) addFacility(name string, numHold, amount float64) {
	tmp := facility.Generate(len(data.facilities), name, numHold, amount)
	data.facilities[name] = &tmp
	tmp.CountUp(data)
}

// InitFacility is init facilities and exec count up
func (data *Data) InitFacility() {
	data.addFacility("student", 1, 100)
	data.addFacility("3d-printer", 8, 2400)
	data.addFacility("senbei-refining", 108, 66000)
	data.addFacility("router", 1680, 1920000)
	data.addFacility("lightning-rod", 8900, 63900000)
	data.addFacility("compiler", 524288, 2147483647)
	data.addFacility("senbei-chemical-reactor", 7777777, 81900000000)
	data.addFacility("graphic-card", 173000000, 3070308030900)
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
