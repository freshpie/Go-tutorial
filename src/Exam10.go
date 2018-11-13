package main

import "fmt"

func main() {
	result := func(datas ...int) int {
		counter := 0

		for _, data := range datas {
			counter += 1
			fmt.Println(data)
		}
		return counter
	}

	totalData := result(2, 3, 4, 5, 6, 7, 8)

	println(totalData)

	plus := func(i int, j int) int {
		tot := i + j
		return tot
	}

	fmt.Println("=======================================")
	plus2 := func(i int, j int) int {
		tot := i - j
		return tot
	}
	fmt.Println("plus2 address : ", plus2)
	fmt.Println(plus2(100, 50))
	changeFunction(plus2, 100, 200)
	fmt.Println("after plus2 address : ", plus2)
	fmt.Println(plus2(100, 50))

	fmt.Println("=======================================")
	plus3 := func(i2 int, j2 int) int {
		tot := i2 - j2
		return tot
	}
	fmt.Println("plus3 address : ", plus3)
	fmt.Println(plus3(100, 50))
	changePointer(&plus3, 100, 200)
	fmt.Println("after plus3 address : ", plus3)
	fmt.Println(plus3(100, 50))
	fmt.Println("=======================================")

	totalData2 := calc(plus, 100, 200)
	fmt.Println(totalData2)
	totalData3 := calc(func(a int, b int) int { return a * b }, 100, 200)
	fmt.Println(totalData3)
}

func calc(t func(int, int) int, fData int, sData int) int {
	return t(fData, sData)
}

func changeFunction(t func(int, int) int, fData int, sData int) {
	s := func(q int, w int) int {
		return q + w
	}
	fmt.Println("t address : ", t)
	fmt.Println("s address : ", s)
	t = s

	fmt.Println("change t address : ", t)
	fmt.Println("change s address : ", s)
}

func changePointer(t *func(int, int) int, fData int, sData int) {
	s := func(q int, w int) int {
		return q + w
	}
	fmt.Println("t address : ", *t)
	fmt.Println("s address : ", s)
	*t = s

	fmt.Println("change t address : ", *t)
	fmt.Println("change s address : ", s)
}
