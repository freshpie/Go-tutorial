package main

func main() {
	var ar [5]int
	ar[0] = 1
	ar[1] = 2
	ar[2] = 3
	ar[3] = 4
	ar[4] = 5
	var ar1 = [5]int{2, 4, 1, 5, 5}
	var ar2 = [...]int{2, 56, 7, 2, 6}

	for i := 0; i < 5; i++ {
		println(ar[i])
	}

	// Slice (동적(가변) 배열)
	var ar3 []int // 동적 배열 1차원가능(linked list)
	ar3 = []int{321, 3, 16}
	ar3[1] = 1000
	println(ar3) // 321, 1000, 16

	ar4 := make([]int, 10, 20)
	println(len(ar4), cap(ar4))
	// len 내장 함수  길이
	// cap 내장 함수 용량

	ar5 := []int{10, 346, 13, 47, 7, 100}
	ar5 = ar5[2:5] // 부분 슬라이드
	ar5 = ar5[:5]
	ar5 = ar5[1:]
	ar5 = ar5[:]
	println(ar5)
	// 2 ~ 5 사용하면 13, 47, 7 마지막 인덱스 + 1 사용
	ar5 = append(ar5, 3)
	ar5 = append(ar5, 3, 4, 45, 16)

}
