package main

import "fmt"

///интерфейс

type structHare struct {
	N1, N2 int
}

func (s *structHare) Sum() int {
	return s.N1 + s.N2

}

type InterfaceHere interface {
	Sum() int
}

func main() {
	//var a InterfaceHere
	//sh := structHare{1, 2}
	//os := OtherStruct{2, 3}

	//a = &sh
	///fmt.Println(a.Sum())
	//a = os
	///fmt.Println(a.Sum())
	var i InterfaceHere = OtherStruct{3, 3}
	///fmt.Println(i.Sum())
	fmt.Printf("(%v, %T)\n", i, i)

}

type OtherStruct struct {
	A, B int
}

func (o OtherStruct) Sum() int {
	return o.A + o.B

}
