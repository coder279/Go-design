package main

import (
	"fmt"
)

type API interface {
	Test() string
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

type aModuleImpl struct{}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type bModuleImpl struct{}

func (b *bModuleImpl) TestB() string {
	return "B module running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func main() {
	a := NewAPI()
	s := a.Test()
	fmt.Println(s)
}
