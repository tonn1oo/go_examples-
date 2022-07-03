package main

import "fmt"

/// error

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (o *AppError) Error() string {
	fmt.Println(o.Custom)
	fmt.Println(o.Field)
	return fmt.Sprintf("%s error", o.Custom)
}

func (o *AppError) Unwrap() error {
	return o.Err

}

func main() {
	err := method1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")

	//err := method1()
	//if err != nil {
	//	fmt.Println(err.Error())

}

//func method1() error {
//	return &AppError{
//		Err:    fmt.Errorf("my error"),
//		Custom: "value here",
//		Field:  89,
//	}
//}

func method1() error {
	return method2()
}

func method2() error {
	return method3()
}
func method3() error {
	return &AppError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong",
	} ///nil ///fmt.Errorf("some error")
}
