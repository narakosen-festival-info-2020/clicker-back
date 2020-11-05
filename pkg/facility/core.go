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
	id      int
	name    string
	numHold int
	numGen  float64
	amount  float64
	sync.RWMutex
}

// Generate facility
func Generate(id int, name string, numGen, amount float64) Data {
	return Data{
		id:      id,
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
			click.AddCount((float64)(data.numHold) * data.numGen / 10)
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
	return (float64)(data.numHold) * data.numGen
}

// GetNumHold is return numHold
func (data *Data) GetNumHold() int {
	return data.numHold
}

// UpgradeByAchieve is applly all achieve
func (data *Data) UpgradeByAchieve() {
	// nothing
}

// UpgradeByInherentAchieve is apply inherent achieve
func (data *Data) UpgradeByInherentAchieve() {
	data.Lock()
	defer data.Unlock()
	data.numGen *= 1.5
}
