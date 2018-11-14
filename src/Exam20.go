package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson() *person {
	p := person{}
	return &p
}

func main() {
	p1 := newPerson()
	p1.name = "leehan"
	p1.age = 18

	p2 := newPerson()
	p2.name = "jwon"
	p2.age = 20

	fmt.Println(p1)
	fmt.Println(p2)
}
