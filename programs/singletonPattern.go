package programs

import "sync"

var lock = &sync.Mutex{}
type Mystruct struct{
	name string
	age int
}

var singleInstance *Mystruct

func getInstance() *Mystruct{
	if singleInstance == nil{
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &Mystruct{name:"sagar", age:25}
	}
	return singleInstance
}
