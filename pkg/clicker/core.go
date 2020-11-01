package clicker

import (
	"strconv"
	"sync"
	"time"

	"github.com/narakosen-festival-info-2020/clicker-back/pkg/achieve"
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/status"
)

// Data is clicker count data :)
type Data struct {
	totalCount    float64
	count         float64
	clickCount    float64
	clickGenCount float64
	perClick      float64
	achieveCount  int
	facilities    map[string]*facility.Data // not under sync (sync only facility.Data)
	statements    status.Data               // not under sync (sync only status.Data)
	achievements  achieve.Data              // not under sync (sync only achieve.Data)
	sync.RWMutex
}

// Generate Clicker
func Generate() *Data {
	ret := Data{
		totalCount:    0,
		count:         0,
		clickCount:    0,
		clickGenCount: 0,
		perClick:      1,
		achieveCount:  0,
		facilities:    make(map[string]*facility.Data),
		statements:    status.Data{},
		achievements:  achieve.Data{},
	}
	update := func() {
		for {
			ret.perClick = 1 + ret.GetSembeiPerSecond()/100*(float64)(ret.achieveCount)
			time.Sleep(time.Second / 100)
		}
	}
	go update()
	return &ret
}

// AddCount is plus count
func (data *Data) AddCount(cnt float64) {
	data.Lock()
	defer data.Unlock()

	data.count += cnt
	data.totalCount += cnt
}

// AddClick is plus click count
func (data *Data) AddClick(cnt float64) {
	data.Lock()
	defer data.Unlock()
	tmp := cnt * data.perClick
	if cnt <= 0 {
		return
	}
	if cnt > 20 {
		tmp = 20 * data.perClick
	}
	data.count += tmp
	data.totalCount += tmp
	data.clickCount += cnt
	data.clickGenCount += tmp
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

// GetSembeiPerSecond is get SpS (all facility SpS)
func (data *Data) GetSembeiPerSecond() float64 {
	ret := 0.0
	for _, tmp := range data.facilities {
		ret += tmp.GetProductionEfficiency()
	}
	return ret
}

// UpgradeByAchieve is applly all achieve
func (data *Data) UpgradeByAchieve() {
	// nothing
}

// UpgradeByInherentAchieve is apply inherent achieve
func (data *Data) UpgradeByInherentAchieve() {
	data.Lock()
	defer data.Unlock()
	data.achieveCount++
}

// InitStatements is init statements(all status) and automatic update
func (data *Data) InitStatements(otherGeneral, otherClick map[string]func() float64) {
	defer data.statements.InitUpdate()

	data.statements.AddGeneral("total_count", func() float64 {
		data.RLock()
		defer data.RUnlock()
		return data.totalCount
	})
	data.statements.AddGeneral("now_count", func() float64 {
		data.RLock()
		defer data.RUnlock()
		return data.count
	})
	// now Senbei per Second
	data.statements.AddGeneral("now_sps", func() float64 {
		return data.GetSembeiPerSecond()
	})

	// total click
	data.statements.AddClick("total_click", func() float64 {
		data.RLock()
		defer data.RUnlock()
		return data.clickCount
	})

	data.statements.AddClick("total_click_gen", func() float64 {
		data.RLock()
		defer data.RUnlock()
		return data.clickGenCount
	})

	data.statements.AddClick("now_spc", func() float64 {
		data.RLock()
		defer data.RUnlock()
		return data.perClick
	})
	for name, tmp := range otherGeneral {
		data.statements.AddGeneral(name, tmp)
	}
	for name, tmp := range otherClick {
		data.statements.AddClick(name, tmp)
	}
}

// InitAchivements is init achivements(all achieve) and automatic check
func (data *Data) InitAchivements(other map[string]func() bool, otherName []string, otherInherent map[string][]achieve.Upgrade) {
	defer data.achievements.InitCheck()

	general := make([]achieve.Upgrade, 0)
	for _, tmp := range data.facilities {
		general = append(general, tmp)
	}

	// about clicker
	for idx, value := 0, 1.0; idx <= 150; idx += 3 {
		tmp := value
		data.achievements.Add("total-senbei-"+strconv.Itoa(idx), func() bool {
			return data.totalCount >= tmp
		}, general, []achieve.Upgrade{})
		value *= 1000
	}

	// about sps
	for idx, value := 0, 1.0; idx <= 150; idx += 3 {
		tmp := value
		data.achievements.Add("sps-"+strconv.Itoa(idx), func() bool {
			return data.GetSembeiPerSecond() >= tmp
		}, general, []achieve.Upgrade{})
		value *= 1000
	}

	// about gen click
	for idx, value := 0, 1.0; idx <= 150; idx += 3 {
		tmp := value
		data.achievements.Add("click-senbei-"+strconv.Itoa(idx), func() bool {
			return data.clickGenCount >= tmp
		}, general, []achieve.Upgrade{})
		value *= 1000
	}

	numSuccess := []int{
		1, 5, 10, 50, 100, 150, 200, 250, 300, 400, 600, 750, 1200, 1500,
	}

	for name, tmp := range data.facilities {
		for _, value := range numSuccess {
			tmpV := value
			tmpF := tmp
			data.achievements.Add(name+"-"+strconv.Itoa(value), func() bool {
				return tmpF.GetNumHold() >= tmpV
			}, general, []achieve.Upgrade{
				tmpF,
			})
		}
	}

	for _, tmp := range otherName {
		data.achievements.Add(tmp, other[tmp], general, otherInherent[tmp])
	}
}
