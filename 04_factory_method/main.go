package main

import "log"

// factory_method 工厂方法其实是子类继承父类去实现的
// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

type MinusOperatorFactory struct {
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func main() {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		log.Println("calc failed !!")
	}
	factory = MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		log.Println("calc failed !!")
	}
}
