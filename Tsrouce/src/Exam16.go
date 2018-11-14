package main

type Rect struct {
	width, height int
}

// 메서드 <- value recevier
// Go 구조체의 필드를 이용해서 기능을 만든 것
//  메서드(구조체를 이용한 함수 만들기)
func (r Rect) area() int {
	return r.width * r.height
}

//value receviver : 원래 데이터를 복사하고 전달하는데
//pointer recevier - 포인터만 전달
// 메서드 내에 필드 값이 변경이 그대로 호출하는 곳에 전달
// 되는 것
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()   // 메서드 호출
	area2 := rect.area2() // 메서드 호출
	println(area)
	println(rect.width, area2)

}
