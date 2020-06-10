package main

// c 언어에서 stdio.h 와 같다고 생각하면 됩니다.
import "fmt"

func main(){
    //반목문
    //go 언어에서는 java, php, c++, javascript 와 같이  for, foreach 등 여러가지 for 문이 있지만
    // 파이썬(while 있음)과 비슷하게 반복문은  for문 하나 있습니다.

    i := 0
    // go 언어는 조건문에 괄호는 없습니다.
    for i <=10{
        fmt.Println(i)
        i += 1
    }
    for i:=10; i>=0; i--{
        fmt.Println(i)
    }
    /*
        break, continue 는 다른언어와 똑같습니다.
        break문을 많이 사용하는것은 좋지 않을 수도 있습니다.

    */

    var array [10]int
    // go 언어는 사용하지 않는 변수를 선언 할 시 오류를 발생 시킵니다.
    // 안쓰는 변순는 _사용하면 문제가 해결 됩니다.
    for i, _  := range array{
        fmt.Println(i)
    }


}