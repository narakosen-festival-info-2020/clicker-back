package facility

import (
	"sync"
	"time"
)

// Click is facility must use func
type Click interface {
	AddCount(float64)
	MinusCount(float64) bool
}

// Data is facility Data
type Data struct {
	name    string
	numHold float64
	numGen  float64
	amount  float64
	sync.RWMutex
}

// Generate facility
func Generate(name string, numGen, amount float64) Data {
	return Data{
		name:    name,
		numHold: 0,
		numGen:  numGen,
		amount:  amount,
	}
}

// CountUp is Generating an infinite number of cookies
func (data *Data) CountUp(click Click) {
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			data.RLock()
			click.AddCount(data.numHold * data.numGen / 10)
			data.RUnlock()
		}
	}()
}

// Purchase is numHold increment
func (data *Data) Purchase(click Click) bool {
	if !click.MinusCount(data.amount) {
		return false
	}
	data.Lock()
	defer data.Unlock()
	data.numHold++
	data.amount *= 1.1
	return true
}

// GetProductionEfficiency is return Sembei per Second
func (data *Data) GetProductionEfficiency() float64 {
	data.RLock()
	defer data.RUnlock()
	return data.numHold * data.numGen
}
