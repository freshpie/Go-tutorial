package main

import "os"

func main() {
  file, err := os.Open("test.txt")
  if err != nil {
    panic(err)
  }

  // 지연 실행 : 정상적일 또는 오류일 때 처리하는 것
  defer file.Close()

  openFile("test.txt")
  println("처리 끝")
}

func openFile(f string) {
  defer func () {
    if r:= recover(); r != nil {
      println("열기 오류", r)
    }
  }

  file, err := os.Open(f)
  if err != nil {
    panic(err)
  }

  defer file.Close()
}
