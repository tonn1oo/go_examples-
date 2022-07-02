package main

/// error

type error interface {
	Error() string
}

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func main() {

}
