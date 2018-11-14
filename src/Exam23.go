package main

import (
	"fmt"
	"time"
)

func sum(num int) {
	var sum int = 0
	var i int
	for i = 0; i <= num; i++ {
		sum += i
	}
	println("sum : ", sum)
}
func sum2(num int, result chan bool) {
	var sum int = 0
	var i int
	for i = 0; i <= num; i++ {
		sum += i
	}
	println("sum2 : ", sum)
	result <- true

}
func main() {
	//runtime.GOMAXPROCS(8)
	fmt.Println(time.Now().Format(time.RFC850))
	go sum(100000000)
	go sum(100000000)
	go sum(100000000)
	go sum(100000000)
	//time.Sleep(time.Second * 5)
	fmt.Println(time.Now().Format(time.RFC850))

	/*ch := make(chan bool)
	go sum2(10000000, ch)
	x := <- ch
	println(x)*/

	ch2 := make(chan bool)
	go sum(10000000)
	for i := range ch2 {
		println(i)
	}

}
