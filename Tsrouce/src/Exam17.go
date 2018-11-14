package main

type ImyInof interface {
	area() float32
	mydata()
}

type Rect struct {
	w, h float32
}

type myInfo struct {
	name string
	age  int
}

func (r Rect) area() float32 {
	return r.w * r.h
}

func main() {
	r := Rect{10, 20}
	result := area()

	var iData interface{} = 1
	iData2 := iData          // iData / iData2 동적 타입
	iData = "홍길동"            // 가능 왜? 동적 타입
	iData3 := iData.(string) // 타입을 확인

}
