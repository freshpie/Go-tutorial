package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

type MyError struct {
	Msg string
	Err error
}

func main() {
	f, err := os.Open("C:\\Go\\README.md")
	if err != nil {
		log.Panic(err.Error())
	}
	println(f.Name())

	new(MyError{"awef", syscall.EINVAL})
	calc(3, 5)

	type person struct {
		name  string
		age   int
		phone string
	}

	leehan := person{name: "이한", age: 30, phone: "010-4058-4049"}

	fmt.Println(leehan)
}

func calc(fData int, sData int) (int, error) {
	return (fData + sData), syscall.EINVAL
}
