package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("C:\\GoApp\\FirstPrj\\src\\Hello.go")

	if err != nil {
		log.Fatal("파일이 없습니다.")
	} else {
		println(file.Name())
	}
}
