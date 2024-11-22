package main

import "fmt"

/*
> - 简单工厂：唯一工厂类，一个产品抽象类，工厂类的创建方法依据入参判断并创建具体产品对象。
> - 工厂方法：多个工厂类，一个产品抽象类，利用多态创建不同的产品对象，避免了大量的if-else判断。
> - 抽象工厂：多个工厂类，多个产品抽象类，产品子类分组，同一个工厂实现类创建同组中的不同产品，减少了工厂子类的数量。
*/

// Operator 被封装的实际接口
type Operator interface {
	SetAB(int, int)
	Result() int
}
// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}
func (o *OperatorBase) SetAB(a int, b int) {
	o.a = a
	o.b = b
}
// PlusOperatorFactory  加法运算的工厂类
type PlusOperator struct {
	*OperatorBase
}
func (p *PlusOperator) Result() int {
	return p.a + p.b
}
type MinusOperator struct {
	*OperatorBase
}
func (p *MinusOperator) Result() int {
	return p.a - p.b
}

type PlusOperatorFactory struct{}
type MinusOperatorFactory struct {}
//加法工厂
func (p *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}
//减法工厂
func (p *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

func main() {
	//加法
	plusFactory := PlusOperatorFactory{}
	plusOperator := plusFactory.Create()
	plusOperator.SetAB(10, 20)
	result := plusOperator.Result()
	fmt.Println("plusOperator=", result)

	//减法
	minusFactory := MinusOperatorFactory{}
	minusOperator := minusFactory.Create()
	minusOperator.SetAB(10, 5)
	result = minusOperator.Result()
	fmt.Println("minusOperator=", result)
}
