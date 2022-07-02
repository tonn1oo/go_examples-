package main

import "fmt"

func main() {

	var myString string
	var hello string = "Hello"
	var title string = "World"

	word := "my string" //автоматическое определение типа

	str := "fhg"
	str = "Hello"

	var b byte = '£' // равно 163
	var r rune = '◯' //9711

	str = "some text"
	fmt.Println(len(str)) //9
	fmt.Println(str[1])   ///111

	str = "строка"
	fmt.Println(len(str))

	fmt.Println(myString, title, hello, word, str, b, r)

	s := "Privet"
	s = s + "!!!!!"
	fmt.Println(s)
}
