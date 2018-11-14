package main

import "fmt"

//Closure

func nextA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var ht map[int]string
	ht = make(map[int]string)

	ht[1] = "KTDS"
	ht[2] = "KT"

	println(ht[1])
	println(ht[100])
	println(ht[2])

	//delete(ht, 1)
	println(ht[1])
	println(ht[2])

	key1 := map[string]string{
		"ktds": "ktds inc",
		"kt":   "kt inc",
	}
	value, exist := key1["ktds"]
	if !exist {
		fmt.Println(value, exist, "없다")
	} else {
		fmt.Println(value, exist, "있다")
	}

	a := key1["ktds"]
	if a != "ktds inc" {
		fmt.Println("없다")
	} else {
		fmt.Println("있다")
	}

}
