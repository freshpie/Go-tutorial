package main

//구조체는 필드만 정의할 수 있다.
type myInfo struct {
	name string
	age  int
}

type dic struct {
	data map[int]string
}

// 생성자 함수
func newDic() *dic {
	myD := dic{}
	myD.data = map[int]string{}
	return &myD // 포인터로 전달
}

func main() {
	info := myInfo{}
	info.name = "홍길동"
	info.age = 20

	println(info.age)

	//var info2 myInfo
	//info2 = myInfo{"홍길동2", 21}
	//info3 := myInfo{name: "홍길동3", age: 22}

	info4 := newDic() // 생성자 호출(생성자 함수 호출)
	info4.data[1] = "KTDS"

	println(info4.data[1])
}
