package main

import "fmt"

// Goods 构建的对象
type Goods struct {
	Name  string
	Price float64
}

// GoodsBuilder 构建器
type GoodsBuilder interface {
	SetName(name string) GoodsBuilder
	SetPrice(price float64) GoodsBuilder
	Build() *Goods
}

// ConcreteBuilder 具体构建器
type ConcreteBuilder struct {
	goods *Goods
}

func (g ConcreteBuilder) Build() *Goods {
	return g.goods
}

func (g ConcreteBuilder) SetName(name string) GoodsBuilder {
	g.goods.Name = name
	return g
}

func (g ConcreteBuilder) SetPrice(price float64) GoodsBuilder {
	g.goods.Price = price
	return g
}

func NewGoodsBuilder() GoodsBuilder {
	return &ConcreteBuilder{
		goods: &Goods{},
	}
}

func main() {
	builder := NewGoodsBuilder()
	goods := builder.SetName("apple").SetPrice(65.0).Build()
	fmt.Println(goods)
}
