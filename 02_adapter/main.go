package main

import "log"

// Target 目标方法里面去适配另一个方法
type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct {
}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}

var expect = "adaptee method"

func main() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		log.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
