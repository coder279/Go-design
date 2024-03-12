package main

import (
	"fmt"
	"log"
)

// API 工厂模式就是将所有的接口都声明好，实现更好的灵活性去适配不同接口的同一个方法
type API interface {
	Say(name string) string
}

func NewAPI(point string) API {
	switch point {
	case "old":
		return &oldAPI{}
	case "new":
		return &newAPI{}
	default:
		return &oldAPI{}
	}
}

type oldAPI struct {
}
type newAPI struct {
}

func (o *oldAPI) Say(name string) string {
	return fmt.Sprintf("Old api %s", name)
}

func (n *newAPI) Say(name string) string {
	return fmt.Sprintf("New api %s", name)
}

func main() {
	api := NewAPI("old")
	content := api.Say("old")
	log.Println(content)
}
