package clicker

import (
	"fmt"
	"sync"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
)

// Data is clicker count data :)
type Data struct {
	count      float64
	facilities map[string]*facility.Data
	sync.RWMutex
}

// Generate Clicker
func Generate() Data {
	return Data{
		count:      0,
		facilities: make(map[string]*facility.Data),
	}
}

// AddCount is plus count
func (data *Data) AddCount(cnt float64) {
	data.Lock()
	defer data.Unlock()
	fmt.Println(data.count)

	data.count += cnt
}

// MinusCount is minus count (etc purchase)
func (data *Data) MinusCount(cnt float64) bool {
	data.Lock()
	defer data.Unlock()
	if data.count < cnt {
		return false
	}
	data.count -= cnt
	return true
}

// GetCount is return click count
func (data *Data) GetCount() float64 {
	data.RLock()
	defer data.RUnlock()
	return data.count
}
