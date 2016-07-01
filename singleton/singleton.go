package singleton

import "sync"

var once sync.Once

type MySingleton struct {
	SomeVariable string
}

var myInstance *MySingleton

func GetMyInstance() *MySingleton {
	once.Do(func() {
		myInstance = &MySingleton{SomeVariable:"Var"}
	})
	return myInstance
}
