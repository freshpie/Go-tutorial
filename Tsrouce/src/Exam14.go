package main

import "Lab/testlib"

// 패키지 : 함수 , 구조체, 인터페이스, 메서드
// 객체지향에서 정보은닉
// 패키지 안에 정보은닉 public 선언하는것
// go 에서는 정보은닉 public / private
// public : 이름을 대문자로 시작하면 무조건 public
// private : 이름을 소문자로 시작하면 무조건 private
func main() {
	mypkg := testlib.GetCompany("GO")
	println(mypkg)
}
