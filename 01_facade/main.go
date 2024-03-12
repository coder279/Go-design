package main

import (
	"fmt"
	"log"
)

// API 外观模式就是对细节的隐藏 内部其实很臃肿 外部暴露一个端口可以做完所有的事情
type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleApi(),
	}
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.Test()
	bRet := a.b.Test()
	return fmt.Sprintf("%s %s", aRet, bRet)
}

type AModuleAPI interface {
	Test() string
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type aModuleImpl struct {
}

func (a *aModuleImpl) Test() string {
	return "A Module running"
}

type BModuleAPI interface {
	Test() string
}

func NewBModuleApi() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct {
}

func (b *bModuleImpl) Test() string {
	return "B Module running"
}

func main() {
	api := NewAPI()
	result := api.Test()
	log.Println(result)
}
