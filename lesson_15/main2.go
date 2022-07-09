package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

//gJason

func main() {
	json := `{
		"name": {
		"first": "Tom", "last": "Anderson"},
		"age":37,
		"children": ["Sara", "Alex", "Jack"],
        "fav.movie": "Deer Hunter",
		"friends": [
{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
]
}`

	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)

		} else if arg == "lower" {
			return strings.ToLower(json)

		}
		return json

	})

	//value := gjson.Get(json, "fav\\.movie")
	//value := gjson.Get(json, "children.2")
	//value := gjson.Get(json, "children.#")
	//value := gjson.Get(json, "child*.2")
	//value := gjson.Get(json, "friends.0.nets")
	//value := gjson.Get(json, `friends.#(age == 47).first`)
	//value := gjson.Get(json, `friends.#(age >= 42)#.last`)

	/// Modifier

	//value := gjson.Get(json, `friends.#(age >= 42)#.last|@case:upper`) // с большой
	value := gjson.Get(json, `friends.#(age >= 42)#.last|@case:lower`) // с маленькой
	fmt.Println(value.String())

	fmt.Println(gjson.Parse(json).Get("name")) // парсить

}
