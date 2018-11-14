package main

//Closure

func nextA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	funcA := nextA()
	println(funcA())
	println(funcA())
	println(funcA())
	println(funcA())

	funcA2 := nextA()
	println(funcA2())
	println(funcA2())
	println(funcA2())

}
