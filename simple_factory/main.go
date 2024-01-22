package main

import (
	"log"
	"strings"
)

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else {
		return &helloAPI{}
	}
}

// hiAPI is one of API implement
type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	var builder strings.Builder
	builder.Write([]byte(name))
	return builder.String()
}

// helloAPI is another API implement
type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	var builder strings.Builder
	builder.Write([]byte(name))
	return builder.String()
}

func main() {
	s := NewAPI(2)
	name := s.Say("hello")
	log.Println(name)
}
