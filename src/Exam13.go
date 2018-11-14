package main

import "fmt"

func main() {

	var yArr1 [100]interface{}
	for index, _ := range yArr1 {
		yArr1[index] = index + 1000
	}

	var yArr [100]interface{}
	for index, _ := range yArr {
		yArr[index] = index + 1000
	}

	for index, _ := range yArr {
		fmt.Printf("11 index : %d\n", yArr[index])
		println("33 index : ", index, yArr[index])
	}
	for index, _ := range yArr {
		fmt.Printf("22 index : %d\n", yArr[index])
		println("44 index : ", index, yArr[index])

	}

}
