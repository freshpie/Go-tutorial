package main

import "fmt"

func main() {
	var x1, x2, x3 interface{}
	x1 = "a"
	x2 = "a"
	x3 = "a"
	println(x1)
	println(x2)
	println(x3)

	var x interface{}

	x = 1
	x = "tom"
	x = 202

	fmt.Printf("%d\n", x)
	println("==================")

	var y interface{}
	println(y)

	println("==================")

	var z1 interface{}
	z1 = "asdf"

	var z2 = z1
	z2 = int64(111)

	var z3 = z1
	z3 = int64(5545555555555555555)

	println("z1 : ", z1)
	println("z2 : ", z2)
	println("z3 : ", z3)

	var i8 int8 = 127
	println("i8 : ", i8)
	var ui8 uint8 = 255
	println("ui8 : ", ui8)
	var i32 int32 = (4294967295 / 2)
	println("i32 : ", i32)

	var ino int = 4294967295000000
	println("ino : ", ino)

}
