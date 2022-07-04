package main

import (
	"fmt"
	"time"
)

//// don't PANIC ! паники в го

type name struct {
	A, B int
}

func (n *name) method() { // создаем метод
	fmt.Println("ok") //метод что то делает
	//fmt.Println(n.A)

}

func main() {

	panic("DON'T PANIC")
	n := &name{A: 1, B: 2} // делаем ссылкой на структуру
	n = nil                // обнуляем
	n.method()
	go test()
	time.Sleep(300 * time.Millisecond)

}

func test() int {
	a := []int{1, 2, 3}
	return a[2]

}
