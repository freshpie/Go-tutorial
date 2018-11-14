package main

import (
	"runtime"
	"time"
)

func myTest(msg string) {
	for i := 0; i < 10; i++ {
		println(msg, " - ", i)
	}
}

func main() {
	myTest("동시성테스트")

	runtime.GOMAXPROCS(4) // 멀티 코어를 사용하는 방8
	go myTest("동시성1")
	go myTest("동시성2")
	go myTest("동시성3")
	go myTest("동시성4")

	time.Sleep(time.Second * 5)

	//time.Sleep(1000 * 5)

}
