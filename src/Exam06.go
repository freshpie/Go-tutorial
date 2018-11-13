package main

func main() {
	var name string = "leehan"
	var div int = 1

	switch div {
	case 1:
		println(name, div)
	case 2:
		println(name, div)
	case 3:
		println(name, div)
	}

	sum := func(nums ...int) int {
		s := 0
		for _, num := range nums {
			s += num
		}
		return s
	}

	println(sum(1, 2, 3, 4, 5))

	var _ string = "leehan"

}
