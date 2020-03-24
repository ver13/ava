package uuid

import (
	"sync"
)

var (
	f *UUID
	
	once sync.Once
)

func init() {
	once.Do(func() {
		f = &UUID{}
	})
}

func GetInstance() *UUID {
	return f
}
