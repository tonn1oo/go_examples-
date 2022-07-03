package main

import "fmt"

func main() {
	var a interface{}
	a = "jelly"
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n\n", a, a)
	a = 44
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n\n", a, a)

	var b interface{} = 3.14
	//s := b.(string)
	//fmt.Println(s)
	//
	//s, ok := b.(string) ////// NO!
	//fmt.Println(s, ok)
	//
	//f, ok := b.(float32)
	//fmt.Println(f, ok)
	//
	switch b.(type) {
	case int:
		fmt.Println("b is int")
	case string:
		fmt.Println("b is string")
	case bool:
		fmt.Println("b is bool")
	case float64:
		fmt.Println("b is float")
	default:
		fmt.Printf("unknown type %T\n", b)
	}
}
