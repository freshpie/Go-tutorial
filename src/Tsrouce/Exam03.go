package main

func main() {
  total := 100 // 100은 실제 값 = 리터럴
  nameTest := '내 이름은 \n
    홍길동 \n
    입니다'
  nameTest2 := "내 이름은 \n 홍길동" + " \n 입니다."

  var exData1 int = 100
  // 형변환은 명시적으로 지정해야함
  var exData2 float32 = float32(exData1)
  var exData3 string
  exData3 = strig(exData1)
  var exData4 string = "1000"
  var exData5 int = int(exData4)
  var exData6 string ="나이는"
  var exData7 int = int(exData6)



}
/* nameTest
 내 이름은 \n
 홍길동 \n
 입니다
nameTest2
 내 이름은
 홍길동
 입니다.
*/
