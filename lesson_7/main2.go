package main

import (
	"fmt"
	"io"
	"strings"
)

type name struct {
	A int
	B string
}

func (n *name) String() string {
	return n.B

}

func main() {
	r := strings.NewReader("Hello world")

	arr := make([]byte, 4)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d err = %v arr = %v\n", n, err, arr)
		fmt.Printf("arr n bytes: %q", arr[:n])
		if err == io.EOF {
			break

		}
	}

}
