package clicker

import (
	"sync"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/status"
)

// Data is clicker count data :)
type Data struct {
	count      float64
	facilities map[string]*facility.Data
	statements status.Data // not under sync (sync only status.Data)
	sync.RWMutex
}

func (data *Data) initStatements() {
	defer data.statements.InitUpdate()
	data.statements.AddGeneral("total_count", func() float64 {
		data.RLock()
		defer data.Unlock()
		return data.count
	})
}

// Generate Clicker
func Generate() *Data {
	ret := Data{
		count:      0,
		facilities: make(map[string]*facility.Data),
		statements: status.Data{},
	}
	ret.initStatements()
	return &ret
}

// AddCount is plus count
func (data *Data) AddCount(cnt float64) {
	data.Lock()
	defer data.Unlock()

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
