package main
/*
 java, c, c++ switch와 사용법이 같다.
*/
import "fmt"

func main(){

    var i int

    fmt.Scanf("%d",&i)

    switch i{
      case 1:
        fmt.Println("일")
        // break 문이 없어도 됩니다.
        // 다음 조건문을 통과 시키고 싶으면
        fallthrough
      case 2:
        fmt.Println("둘")
    }

    switch{
      // 이런 식으로 조건식을 넣을 수 도 있습니다.
      case i >=1 && i %2 ==0 :
          fmt.Println("짝수")

      case i >=1 && i %2 !=0 :
          fmt.Println("홀수.")

    }

}