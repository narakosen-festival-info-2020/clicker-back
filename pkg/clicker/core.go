package clicker

import (
	"fmt"
	"sync"
)

// Data is clicker count data :)
type Data struct {
	count int
	sync.RWMutex
}

func (data *Data) AddCount(cnt int) {
	data.Lock()
	defer data.Unlock()
	fmt.Println(data.count)

	data.count += cnt
}
