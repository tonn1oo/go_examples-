package main

import (
	"encoding/json"
	"fmt"
)

// Json

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"isBlocked"`

	Books []struct {
		Name string `json:"name"`
		Year int    `json:"year"`
	} `json:"books"`
}

func main() {
	bv := User{
		Name:      "Anton",
		Age:       35,
		IsBlocked: true,
	}
	sv := map[string]interface{}{"field": "value", "field1": 123, "field2": true, "arr": []string{"a", "b", "c"}} // json обьект
	//sv := map[string]string{"field": "value", "field1": "123", "field2": "true"} // json обьект
	//sv := []string{"A", "B", "C"}
	strU, _ := json.Marshal(bv)
	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(strU))
	fmt.Println(string(boolVar))

	//json.Unmarshal()

}
