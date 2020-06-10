package main

// c 언어에서 stdio.h 와 같다고 생각하면 됩니다.
import "fmt"

func main(){
    // 조건문
    // go언어에도 goto 문법이 존재하지만 최대한 쓰지 않는 것이 좋습니다.
    // 정말 필요할때 제한적으로 쓰여야만 합니다.

    for i:=0; i<=10; i++{
        // java, c , c++ 언어와 똑같습니다.
        if i%2 == 0 && i != 0{
            fmt.Println("짝수 입니다.", i)
        //else if를 보여 주기 위해서 불 필요한 조건문 생성 했습니다.
        } else if i%2 !=0 && i != 0{
            fmt.Println("홀수 입니다.", i)
        }
    }
}