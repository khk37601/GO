package main

// c 언어에서 stdio.h 와 같다고 생각하면 됩니다.
import "fmt"

func main(){
    //변수 선언 방법
    // go언어는 java, c, c++등과 같이 변수를 지정하지만 Kotlin 변수 선언과 비슷하다고 생각 하면 됩니다.
    // kotlin 변수 선언 var str :String = "hello", var str ="Hello"

    // 명시저 선언
    var str string = "안녕하세요 저는 go언어 입니다"
    fmt.Println(str)
    // 묵시적 선언  컴파일러가 타입을 유추합니다.
    var str2 = "저는 아무개 입니다."
    fmt.Println(str2)

    // 여러개 변수 선언 방법
    var x, y = 20, 520
    var num, test = 10, "hello"
    /* 변수 재선언을 오류를 발생 시킵니다.
    var (
        x = 40
        y = 30
    )
    */
    fmt.Println(x, y, num, test)

    /*
     go언어는 c, c++ 비슷하게 자료형 타입이 int32 int64, float32 float64, uint32... 등이 쓰입니다.
     기본적으로 int, float 는 64비트로 인식 합니다.
    */
    var t int64 = 12343254
    fmt.Println(t)

    // 변수를 초기화 할떄 많이 쓰이는 방법
    s := 0
    fmt.Println(s)
    // 이렇게 사용이 가능합니다.
    for s <= 10{
        s += 1
        fmt.Println(s)
    }

    // 상수 const "(c언어, javascript 와 똑같다고 생각 하면 됩니다.)"
    const str = "KJS"
    fmt.Println(str)
    // 값을 바꾸면 오류 발생 합니다.
    //str = "Good"


}