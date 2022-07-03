package main

import "fmt"

/// Функции и параметры

type Point struct { //структура
	X, Y int
}

func (p Point) movePoint(x, y int) Point {
	p.X += x ///так почти никто не пишет
	p.Y += y
	return p

}
func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y

}

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum

	}

}

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 2)
	fmt.Println(s)
	return result
}

func main() {

	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(p.movePoint(1, 1))
	fmt.Println(p)
	p.movePointPtr(1, 1)
	fmt.Println(p)

	///замыкание

	orderPrice := totalPrice(1) // sum =1
	fmt.Println(orderPrice(1))  // sum = 1+1 = 2
	fmt.Println(orderPrice(10)) // 2+10

	sunCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	result := doSomething(sunCallback, "plus")
	fmt.Println(result)

	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}
	result = doSomething(multipleCallback, "multiple")
	fmt.Println(result)
}
