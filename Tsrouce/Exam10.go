package Exam10

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	println(sum)

	// for 문중에 조건식만 있는 것
	//  while 구문을 대체할때 사용
	// while 반목문 : 몇번 반복할 지 모를때
	i := 1
	for i < 100 {
		// 100 이 넘는지 아닌지 확인
		i *= 2
	}

	for {
		// 무한루프
	}

	nameIndex := []string{"홍길동1", "홍길동2", "홍길동3"}

	for index, nams := range nameIndex {
		println(index, names)
	}

	var temp int = 1

	for temp < 100 {
		if temp == 10 {
			temp += temp
			continue
		}
		temp++
		if temp == 20 {
			break // 반복문을  빠져나오는 것
		}

		if temp == 99 {
			goto END // 레이블 지정 한 곳(이름)으로 이동
		}
	}
END: // 레이블 지정된곳
	println("프로그램 종료")

}
