package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// работа с файлами и выполнение shell комманд

func main() {
	appendToFile()
	readFile()
	writeToFile()
}

func readFile() {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func writeToFile() {
	data := []byte("Anton")
	err := ioutil.WriteFile("text.txt", data, 0600)
	if err != nil {
		fmt.Println(err)
	}
}

func appendToFile() {
	for i := 0; i < 100; i++ {
		f, err := os.OpenFile("text.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		if _, err = f.WriteString(" Tonnioo\n"); err != nil {
			panic(err)
		}

	}
}
