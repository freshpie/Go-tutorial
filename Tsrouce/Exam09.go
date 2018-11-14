package Exam09

func main() {
	result := func(datas ...int) int { // 익명함수
		counter := 0
		//for 반복문 중에 순차적으로 처리하는 반복문
		// 순차는 현재 있는 데이터 숫자만큼 반복
		// for 인덱스, 별칭 := range 원본데이터 {
		// 반복할 작업
		//}
		for _, data := range datas {
			counter += 1
		}
		return counter
	}

	totalData := result(3, 6, 1, 7) // 익명함수 호출
	println(totalData)

	plus := func(i int, j int) int { //익명함수 정의
		tot := i + j
		return tot
	}

	totalData2 := calc(plus, 100, 200)
	println(totalData2)
	totalData3 := calc(func(a int, b int) int { return a * b }, 10, 20)

}

// 일급함수
// t 는 익명함수의 별칭이라고 생각하면 됨
func calc(t func(int, int) int, fData int, sData int) int {
	extFunc := t(fData, sData)
	return extFunc
}
