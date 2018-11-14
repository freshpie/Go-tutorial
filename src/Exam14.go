package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(num int64, c chan int64) {
	var sum int64 = 0
	var i int64
	for i = 0; i <= num; i++ {
		sum += i
	}
	c <- sum
	//fmt.Println(sum)
}
func main() {
	c := make(chan int64)
	runtime.GOMAXPROCS(8)
	fmt.Println(time.Now().Format(time.RFC850))
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(1000000000, c)
	go sum(10000000000, c)
	x1, x2, x3, x4, x5, x6, x7, x8, x9 := <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c

	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)
	fmt.Println(x7)
	fmt.Println(x8)
	fmt.Println(x9)

}
