package main

func main()  {
  ch := make(chan int)  // 정수형 빨대를 하나 생성
  // 채널을 생성함

  go func() {
    ch<-1000  // 채널에 1000
   }

   var data int
   data = <- ch
   println(data)

   ch2 := make(chan bool)
   go func() {
     for i := 0 ; i < 100; i++ {
       println(i)
     }
   }()
   <-ch2 // 앞의 go 루틴이 끝날때 까지 대기

   //채널 버퍼링
   ch3 :=  make(char int, 1)
   ch3 <- 1000 // 가능 : 원래는 되지 않음 그런데 버퍼링 때문에 가능

   println(<-ch3)

   close(ch)
   close(ch2)
   close(ch3)
   /*
   if _, suc = <-; !suc {
     println("없다")
   }

   for {
     if i, suc := <- ch; suc {
       println(i)
     } else {
       break
     }
   }
   */
   for i:= range ch {
     println(i)
   }


}
