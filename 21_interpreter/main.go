package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}


type NumberExpression struct {
	val int
}

func (n *NumberExpression) Interpret() int {
	return n.val
}


type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Interpret() int {
    fmt.Printf("Interpret: %d + %d", a.left.Interpret(), a.right.Interpret())
	return a.left.Interpret() + a.right.Interpret()
}


type SubtractionExpression struct {
	left, right Expression
}

func (a *SubtractionExpression) Interpret() int {
    fmt.Printf("Interpret: %d - %d", a.left.Interpret(), a.right.Interpret())
	return a.left.Interpret() - a.right.Interpret()
}



type Parser struct {
	exp   []string
	index int
	prev  Expression
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAdditionExpression()
		case "-":
			p.prev = p.newSubtractionExpression()
		default:
			p.prev = p.newNumberExpression()
		}
		fmt.Println("prev:",fmt.Sprintf("%v",p));
	}
}

func (p *Parser) newAdditionExpression() Expression {
	p.index++
	return &AdditionExpression{
		left:  p.prev,
		right: p.newNumberExpression(),
	}
}

func (p *Parser) newSubtractionExpression() Expression {
	p.index++
	return &SubtractionExpression{
		left:  p.prev,
		right: p.newNumberExpression(),
	}
}

func (p *Parser) newNumberExpression() Expression {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &NumberExpression{
		val: v,
	}
}

func (p *Parser) Result() Expression {
    fmt.Println("Result:",fmt.Sprintf("%v",p));
	return p.prev
}

func main() {
	p := &Parser{}
	p.Parse("3 + 1 - 2")
	res := p.Result().Interpret()

	fmt.Printf("got: %d", res)
}