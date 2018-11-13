package Exam04

func main() {
	var count int = 1
	var fData int = 100
	var sData int = 200
	var total int
	total = fData + sData
	total = fData - sData
	total = fData * sData
	total = fData / sData
	total = fData % sData
	//코드의 가독성
	//total = fData + sData * count
	total = (fData + sData) * count
	count++
	// count = count + 1
	var result bool = true
	result = fData < sData  // 관계연산자
	result = fData && sData // 논리연산자

	// 참 / 거짓이 아닌 값을 받기
	result2 := fData & sData //  bitwise 관계연산자
	// 할당 연산자
	count += 100 // count = count + 100

	// 포인터 연산자
	var pData int = 100
	var pData2 = &pData // pData 변수의 주소 값을 할당

}
