package status

import (
	"sync"
	"time"
)

// Core is status based data
type Core struct {
	name   string
	value  float64
	update func() float64
	sync.RWMutex
}

func generateCore(name string, update func() float64) Core {
	return Core{
		name:   name,
		value:  0,
		update: update,
	}
}

// Update Core data
func (data *Core) Update() {
	data.Lock()
	defer data.Unlock()
	data.value = data.update()
}

// Data is status data
type Data struct {
	general []*Core
	click   []*Core
}

// AddGeneral is data into general core
func (data *Data) AddGeneral(name string, update func() float64) {
	tmp := generateCore(name, update)
	data.general = append(data.general, &tmp)
}

// AddClick is data into click core
func (data *Data) AddClick(name string, update func() float64) {
	tmp := generateCore(name, update)
	data.click = append(data.click, &tmp)
}

// InitUpdate is loop Update core data
func (data *Data) InitUpdate() {
	times := func() {
		for {
			for _, tmp := range data.general {
				tmp.Update()
			}
			for _, tmp := range data.click {
				tmp.Update()
			}
			time.Sleep(time.Second / 10)
		}
	}
	go times()
}
