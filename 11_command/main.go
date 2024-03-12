package main

import "fmt"

type Command interface {
	Execute()
}

type MethodBoard struct{}

type StartCommand struct {
	mb *MethodBoard
}

func NewStartCommand(mb *MethodBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

type RebootCommand struct {
	mb *MethodBoard
}

func NewRebootCommand(mb *MethodBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}
func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

func (*MethodBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MethodBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
func main() {
	mb := &MethodBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()

}
