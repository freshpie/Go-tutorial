package main

import "fmt"

func main() {
	/*var x [3]interface{}
	//println(x)

	x[0] = "a"
	x[1] = "a"
	x[2] = "a"
	println(x[0])
	println(x[1])
	println(x[2])*/

	/*var yArr1 [100]interface{}
	for index, _ := range yArr1{
		yArr1[index] = index + 1000
	}
	for index, _ := range yArr1{
		println("11 index : ", index, yArr1[index])
	}


	var yArr [100]interface{}
	for index, _ := range yArr{
		yArr[index] = index + 1000
	}
	for index, _ := range yArr{
		fmt.Printf("%d\n", yArr[index])
	}

	fmt.Println(yArr)*/

	/*var yArr1 [100]interface{}
	for index, _ := range yArr1{
		yArr1[index] = index + 1000
	}
	var yArr2 [100]interface{}
	for index, _ := range yArr2{
		yArr2[index] = index + 2000
	}

	for index, _ := range yArr1{
		println("11 index : ", index, yArr1[index])
	}
	for index, _ := range yArr2{
		println("22 index : ", index, yArr2[index])
	}*/
	var b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var a []int
	a = []int{1, 2, 3, 4, 5}
	//a[5] = 6
	//a[4] = 6
	a = a[2:4]

	fmt.Println(a)

	c := append(a, 3, 4, 6, 32, 1)
	fmt.Println(c)

	var d = []interface{}{"de", 4, 342.234, 2 + 4, "efwef"}
	fmt.Println(d)
	d = append(d, 33, "23v", "34df")
	fmt.Println("append --> ", d)

}
