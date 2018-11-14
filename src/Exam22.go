package main

import (
	"os"
)

func main() {
	/*file, err := os.Open("C:\\GoApp\\FirstPrj\\src\\test.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()*/

	openFile("C:\\GoApp\\FirstPrj\\src\\Hello..go")
	println("처리 끝")
}

func openFile(f string) {
	defer func() {
		if r := recover(); r != nil {
			println("열기 오류", r)
		}
	}()

	println("open start")
	file, err := os.Open(f)

	if err != nil {
		println("panic call")
		panic(err)
		println("after panic recover")
	} else {
		println("open file name : ", file.Name())
	}

	defer func() {
		println("close file")
		file.Close()
	}()
}
