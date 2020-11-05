package achieve

import (
	"sync"
	"time"
)

// Upgrade by achieve
type Upgrade interface {
	UpgradeByAchieve()
	UpgradeByInherentAchieve()
}

// Core is achievement based data
type Core struct {
	id       int
	name     string
	unlocked bool
	check    func() bool
	general  []Upgrade
	inherent []Upgrade
	sync.RWMutex
}

func generateCore(id int, name string, check func() bool, general, inherent []Upgrade) Core {
	return Core{
		id:       id,
		name:     name,
		unlocked: false,
		general:  general,
		inherent: inherent,
		check:    check,
	}
}

// Check isUnlock Core data
func (data *Core) Check() {
	if data.unlocked {
		return
	}
	data.Lock()
	defer data.Unlock()
	data.unlocked = data.check()
	if data.unlocked {
		for _, tmp := range data.inherent {
			tmp.UpgradeByInherentAchieve()
		}
		for _, tmp := range data.general {
			tmp.UpgradeByAchieve()
		}
	}
}

// Data is achievements data
type Data struct {
	nowID        int
	achievements []*Core
}

// Add is data into core
func (data *Data) Add(name string, check func() bool, general, inherent []Upgrade) {
	tmp := generateCore(data.nowID, name, check, general, inherent)
	data.nowID++
	data.achievements = append(data.achievements, &tmp)
}

// InitCheck is loop Check core data
func (data *Data) InitCheck() {
	times := func() {
		for {
			for _, tmp := range data.achievements {
				tmp.Check()
			}
			time.Sleep(time.Second / 10)
		}
	}
	go times()
}
