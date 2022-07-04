package main

import (
	"fmt"
	"time"
)

/// go рутины

func main() {
	ch := make(chan string) // канал го рутыны c с типом int
	cp := make(chan int)

	go sayHello(ch)
	s := <-ch
	fmt.Println(s)

	go say("hello hello hello", cp) // создали го рутину
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println(<-cp) //читаем из канала
}

func say(word string, cp chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println(word)
	cp <- 10 // записали в канал тип , в данном случае это 10 , но можно и строку

}

func sayHello(exit chan string) {
	sayHello := string("Golang")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(sayHello)
	}
	exit <- "Top"
}
