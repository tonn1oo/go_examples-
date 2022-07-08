package main

import (
	"fmt"
	"math"
)

///const pi = 3.14

type Shape interface {
	Area() float32
}

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{10}

	printShapeArea(square)
	printShapeArea(circle)
}
func printShapeArea(shape Shape) { // создаем фунцию printShapeArea передаем интерфейс,
	// эта функция принемает интерфейс в качестве арумента
	fmt.Println(shape.Area()) //
}
