package main

import (
	"fmt"
	"math"
)

const weight = 70
const C = 3000000

func main() {
	const enery = weight * C
	fmt.Println(enery)
	b := math.Pow(enery, 2)
	fmt.Println(b)
}
