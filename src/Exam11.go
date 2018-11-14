package main

import "fmt"

func main() {
	type person struct {
		name  string
		age   int
		phone string
	}

	leehan := person{name: "이한", age: 30, phone: "010-4058-4049"}

	fmt.Println(leehan)

}

func calc(t func(int, int) int, fData int, sData int) int {
	return t(fData, sData)
}
