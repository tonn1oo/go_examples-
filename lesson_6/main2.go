package main

import "fmt"

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

func (u *user) RenameFIO(newFio string) {
	u.FIO = newFio
}

func (u *user) RenameAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	RenameFIO(newFio string)
	RenameAddress(newAddress string)
}

func NewUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &u

}

func main() {
	user := NewUser("qwer", "asdfg", "1234556")
	user.RenameFIO("ant")
	user.RenameAddress("arm")

	fmt.Println(user)
}
