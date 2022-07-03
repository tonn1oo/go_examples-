package main

import (
	"fmt"
)

///Указатели, структуры, массивы и слайсы

type Point struct {
	X int
	Y int
	S string
}
type Toni struct { //структура
	rec string
	edc string
	qwe int
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)

}

func main() {
	pointers()
	structs()
	test()

	var array [3]string //массив
	array[0] = "go"
	array[1] = "lang"
	array[2] = "what?"
	fmt.Println(array)
	fmt.Println(array[1])
	fmt.Println(array[2])

	numbers := [...]int{1, 2, 3, 4}
	numbers[2] = 4
	numbers[3] = 5
	fmt.Println(numbers)
	// slice
	letters := []string{"a", "b", "c"}       //создали слайс
	letters[0] = "1"                         // a заменили на 1
	letters = append(letters, "d")           //добавили в слайс d
	letters = append(letters, "e", "f", "g") //добавили в слайс
	fmt.Println(letters)

	createSlice := make([]string, 3) // можно и пустым , без значений//создали слайс размером 3 элемента
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	//createSlice[3] = "4" // error
	createSlice = append(createSlice, "4") //добавили элемент
	fmt.Println(createSlice)

	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}
	v := animalsArr[0:2] /// slice
	fmt.Println(v)
	x := animalsArr[1:4]
	fmt.Println(x)

	x[0] = "123"
	fmt.Println(v)
	fmt.Println(animalsArr)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[:5] //с 0 до 5
	//t := s[0:5] //с 0 до 5 [1, 2, 3, 4, 5]
	fmt.Println(t)
	//tt := s[5:10] // c 5 по 10 [6, 7, 8, 9, 10]
	tt := s[5:]
	fmt.Println(tt)
	ttt := s[0:] //с 0 и до конца
	fmt.Println(ttt)
	tttt := s[:] ///с 0 и до конца
	fmt.Println(tttt)

	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("slice is nil")
	}

}

func pointers() {
	a := "Qwerty"

	p := &a //указатели
	fmt.Println(p)
	fmt.Println(*p)
	*p = "oh my"
	fmt.Println(a)

	b := 42
	g := &b
	*g = *g / 2
	fmt.Println(g)
	fmt.Println(b)
}
func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "hello",
	}
	p1.method()

	fmt.Println(p1)
	fmt.Println(p1.Y)
	p1.X = 1234
	fmt.Println(p1.X)

	p2 := Point{
		Y: 258,
	}
	fmt.Println(p2)
}
func test() {
	a1 := Toni{
		rec: "sunday",
		edc: "monday",
		qwe: 11,
	}
	fmt.Println(a1.rec)
	fmt.Println(a1.edc)
	fmt.Println(a1.qwe)

	g := &a1
	fmt.Println(g)
	fmt.Println(*g)
	fmt.Println(&g)
}
