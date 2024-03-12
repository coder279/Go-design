package main

import (
	"log"
	"sync"
)

type Singleton interface {
	foo()
}
type singleton struct {
}

func (s singleton) foo() {
	log.Println("foo")
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
func main() {
	f := GetInstance()
	f.foo()
}
