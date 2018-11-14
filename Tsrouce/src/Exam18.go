package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("C:\\temp\\test.txt")
	// 'C:\temp\test.txt'
	if err != nil {
		log.Fatal(err.Error())
	}

	println(file.Name())
}
