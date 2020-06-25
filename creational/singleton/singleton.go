package singleton

import (
	"sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = &single{}
		})
	}
	return singleInstance
}

/*

// var lock = &sync.Mutex{}

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}*/
