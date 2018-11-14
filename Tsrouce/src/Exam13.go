package main

func main() {
	// Map => Hash Table (키, 값) : 키를 이용해서
	// 값을 빨리 찾기 위해서 사용하는 것
	var ht map[int]string

	ht = make(map[int]string)
	// 데이터를 추가
	ht[1] = "KTDS"
	ht[2] = "KT"

	data1 := ht[1]
	data2 := ht[1000] //  값이 없으면?? nil / zreo 돌려줌
	println(data1)
	println(data2)

	// 삭제 remove / clear / del / delete
	delete(ht, 1)

	// 키를 확인하는 방법
	key1 := map[string]string{
		"KTDS": "KTDS Inc",
		"KT":   "KT Inc",
		"GO":   "Google Inc",
		"MS":   "Microsoft",
		"FB":   "Facebook",
	}
	//hashtable(map) 키 확인
	ex := key1["GO"]
	if ex != "GO" {
		println("Go 있다")
	}
}
