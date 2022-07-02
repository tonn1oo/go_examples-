package main

/// Итератор for range и структура map

import "fmt"
import "github.com/mitchellh/mapstructure"

type Poind struct {
	X int //`mapstructure:"xx"`
	Y int //`mapstructure:"yy"`
}

func (p Poind) method() {
	fmt.Println("call Point method")

}

func main() {

	pointMap2 := map[string]int{
		"x": 123,
		"y": 2222,
	}

	for k, v := range pointMap2 {
		fmt.Println(k)
		fmt.Println(v)

	}

	p1 := Poind{}
	mapstructure.Decode(pointMap2, &p1)
	fmt.Println(p1)

	pointsMap := map[string]Poind{
		"b": {
			Y: 13,
			X: 15,
		},
	}
	otherMap := map[string]int{}
	otherPointMap := make(map[int]Poind)

	var oneMorePointsMap map[string]Poind
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Poind{
			"c": {
				Y: 12,
				X: 13,
			},
		}
	}
	pointsMap["a"] = Poind{
		X: 1,
		Y: 2,
	}
	otherPointMap[1] = Poind{1, 2}
	fmt.Println(otherMap)
	fmt.Println(otherPointMap[1])
	fmt.Println(otherPointMap)
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	fmt.Println(oneMorePointsMap["c"])
	oneMorePointsMap["a"].method()

	key := "c"
	value, ok := oneMorePointsMap[key] ///проверить есть ли в мапе такой ключ "c"
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s does exist in map\n", key)
		fmt.Println(value)
	}

	arr := []string{"a", "b", "c"}
	for i := range arr {
		fmt.Println(i)
	}

}
