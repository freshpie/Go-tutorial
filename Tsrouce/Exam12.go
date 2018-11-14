package main

// 함수이름 nextV
// nextV는 i 변수의 소유를 함
// nextV는 익명함수를 정의 함
// 클러저 익명함수가 nextV의 i를 접근해서
// 사용하는 것
func nextV() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// n을 호출을 하면 n의 변수에 할당한 후
	// 계속 함수를 호출하면 이 때마다 함수의 값을
	// 계속 유지하는 것 처럼 되는 것
	n := nextV()
	println(n())
	println(n())
	println(n())

	n2 := nextV()
	println(n2())
	println(n2())
	println(n2())
}
