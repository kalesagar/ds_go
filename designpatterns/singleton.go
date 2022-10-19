package designpatterns

import (
	"sync"
)

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock := sync.Mutex{}
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &single{}
		return singleInstance
	}
	return singleInstance
}
