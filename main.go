package main

import (
	"errors"
	"fmt"
)

type myCalc interface {
	add(float64, float64)
	subtract(float64, float64)
	multi(float64, float64)
	div(float64, float64)
}
type Calc struct {
	myCalc
	num1 float64
	num2 float64
}

type Operator string

func (o Operator) add(a, b float64) {
	fmt.Println(a + b)
}
func (o Operator) subtract(a, b float64) {
	fmt.Println(a - b)
}
func (o Operator) multi(a, b float64) {
	fmt.Println(a * b)
}
func (o Operator) div(a, b float64) {
	fmt.Println(a / b)
}
func main() {
	var a, b float64
	var op string
	_, err := fmt.Scan(&a, &op, &b)
	if err != nil {
		panic(errors.New("Введены не верные значения!"))
	}
	o := Operator(op)
	c := Calc{myCalc: o, num1: a, num2: b}
	switch o {
	case "+":
		c.add(c.num1, c.num2)
	case "-":
		c.subtract(c.num1, c.num2)
	case "*":
		c.multi(c.num1, c.num2)
	case "/":
		if c.num2 == 0 {
			panic(errors.New("Нельзя делить на 0"))
		}
		c.div(c.num1, c.num2)
	default:
		panic(errors.New("Введите правильный оператор"))

	}
}
