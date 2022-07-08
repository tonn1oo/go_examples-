package main

import "fmt"

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type Name string

type User struct {
	name   Name
	age    Age
	sex    string
	weight int
	height int
}

type DumbDataBase struct {
	m map[string]string
}

func NewDumpDataBAse() *DumbDataBase {
	return &DumbDataBase{
		m: make(map[string]string),
	}

}

func NewUser(name, sex string, age, weight, height int) User { //// конструктор
	return User{
		name:   Name(name),
		sex:    sex,
		age:    Age(age),
		weight: weight,
		height: height,
	}

}

func main() {
	user1 := NewUser("Vanya", "Male", 12, 180, 200)
	user2 := User{"Vasya", 23, "male", 55, 100}
	user3 := User{"Vasya", 23, "male", 55, 100}
	user4 := User{"Vasya", 23, "male", 55, 100}
	user5 := User{"Vasya", 23, "male", 55, 100}

	fmt.Println(user4.weight)

	fmt.Println(user1.age.isAdult())

	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user3)
	fmt.Printf("%+v\n", user4)
	fmt.Printf("%+v\n", user5)

	user1.printUserInfo() //// метод
	user2.printUserInfo()
}
func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.height, u.weight)

}
