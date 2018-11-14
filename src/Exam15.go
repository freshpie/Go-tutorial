package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(num int, c chan int) {
	var sum int = 0
	var i int
	for i = 0; i <= num; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)
	c5 := make(chan int)
	c6 := make(chan int)
	c7 := make(chan int)
	c8 := make(chan int)

	runtime.GOMAXPROCS(8)
	fmt.Println(time.Now().Format(time.RFC850))
	go sum(1000000000, c1)
	go sum(1000000000, c2)
	go sum(1000000000, c3)
	go sum(1000000000, c4)
	go sum(1000000000, c5)
	go sum(1000000000, c6)
	go sum(1000000000, c7)
	go sum(1000000000, c8)

	i := 0
	for {
		select {
		case x := <-c1:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c1 : ", x)
			i++
		case x := <-c2:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c2 : ", x)
			i++
		case x := <-c3:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c3 : ", x)
			i++
		case x := <-c4:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c4 : ", x)
			i++
		case x := <-c5:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c5 : ", x)
			i++
		case x := <-c6:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c6 : ", x)
			i++
		case x := <-c7:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c7 : ", x)
			i++
		case x := <-c8:
			fmt.Println(time.Now().Format(time.RFC850))
			fmt.Println("c8 : ", x)
			i++
		}
		if i >= 8 {
			break
		}
	}

}
